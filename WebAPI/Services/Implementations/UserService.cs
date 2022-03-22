using AutoMapper;
using Microsoft.Extensions.Logging;
using Models;
using Repositories.Infrastructure;
using Services.Interfaces;
using Services.Models;
using Services.ViewModels;
using System;

namespace Services.Implementations
{
    public class UserService : IUserService
    {
        private readonly IRepository<Users> _userRepository;
        private readonly IUnitOfWork _unitOfWork;
        private ILogger _logger;
        private IMapper _mapper;

        public UserService(IRepository<Users> userRepository,
                           IUnitOfWork unitOfWork,
                           ILoggerFactory loggerFactory,
                           IMapper mapper)
        {
            _userRepository = userRepository;
            _unitOfWork = unitOfWork;
            _logger = loggerFactory.CreateLogger("UserService");
            _mapper = mapper;
        }

        public UserViewModel Authenticate(UserLogin userLogin)
        {
            try
            {
                var result = new UserViewModel();

                var user = _userRepository.FindSingle(_ => _.Username == userLogin.Username);

                if(user != null && BCrypt.Net.BCrypt.Verify(userLogin.Password, user.Password))
                {
                    result = _mapper.Map<UserViewModel>(user);
                    return result;
                }

                return new UserViewModel() { NotFound = true };
            }
            catch(Exception ex)
            {
                _logger.LogError("UserService.Authenticate: " + ex.ToString());
                return null;
            }
        }

        public bool Register(UserRegister userRegister)
        {
            try
            {
                var model = _mapper.Map<Users>(userRegister);

                model.Id = Guid.NewGuid().ToString();
                model.Password = BCrypt.Net.BCrypt.HashPassword(userRegister.Password, BCrypt.Net.BCrypt.GenerateSalt(10));

                _userRepository.Add(model);
                _unitOfWork.Commit();

                return true;
            }
            catch (Exception ex)
            {
                _logger.LogError("UserService.Register: " + ex.ToString());
                return false;
            }
        }
    }
}
