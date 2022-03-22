using Microsoft.Data.SqlClient;
using Microsoft.EntityFrameworkCore;
using Models;
using Moq;
using Repositories.Infrastructure;
using Services.Implementations;
using Services.Interfaces;
using Services.Models;
using Services.ViewModels;
using System;
using System.Linq;
using System.Linq.Expressions;
using System.Text.Json;
using Xunit;

namespace UnitTests
{
    public class UserServiceTest : BaseService_Test, IDisposable
    {
        private IUserService _userService;
        private Mock<IRepository<Users>> _mockUserRepository;

        protected override void Setup()
        {
            _mockUserRepository = new Mock<IRepository<Users>>();

            _mockUserRepository.Setup(x => x.FindSingle(It.IsAny<Expression<Func<Users, bool>>>()))
                               .Returns(new Func<Expression<Func<Users, bool>>, Users>(
                                   (predicate) =>
                                   {
                                       var user = _mockContext.Object.Users;

                                       return user.FirstOrDefault(predicate);
                                   }));

            _mockUserRepository.Setup(x => x.Add(It.IsAny<Users>())).Callback<Users>((u) => 
            {
                if (_mockContext.Object.Users.Any(_ => _.Username == u.Username))
                    throw new Exception("UniqueConstraint");
                _mockContext.Object.Users.Add(u);
            });

            _mockUserRepository.Setup(x => x.DbSet).Returns(_mockContext.Object.Users);

            _mockLoggerFactory.Setup(l => l.CreateLogger("UserService")).Returns(_mockLogger.Object);

            _userService = new UserService(_mockUserRepository.Object, _mockUnitOfWork.Object, _mockLoggerFactory.Object, mapper);
        }

        [Fact]
        public void Should_Register_Success()
        {
            var userRegister = new UserRegister()
            {
                Username = "Member",
                Password = "12345678",
                TaskPerDay = 5
            };

            var result = _userService.Register(userRegister);

            Assert.True(result);
        }

        [Fact]
        public void Should_Register_Exception()
        {
            var userRegister = new UserRegister()
            {
                Username = "Admin",
                Password = "12345678",
                TaskPerDay = 10
            };

            var result = _userService.Register(userRegister);

            Assert.False(result);
        }

        [Fact]
        public void Should_Authenticate_Success()
        {
            var userLogin = new UserLogin()
            {
                Username = "Admin",
                Password = "12345678"
            };

            var userViewModel = _userService.Authenticate(userLogin);

            var result = new UserViewModel()
            {
                ID = "1",
                Username = "Admin",
                NotFound = false
            };

            Assert.NotNull(userViewModel);
            Assert.Equal(JsonSerializer.Serialize(result), JsonSerializer.Serialize(userViewModel));
        }

        [Fact]
        public void Should_Authenticate_UsernameOrPassword_Incorrect()
        {
            var userLogin = new UserLogin()
            {
                Username = "Admin1",
                Password = "12345678"
            };

            var userViewModel = _userService.Authenticate(userLogin);

            Assert.True(userViewModel.NotFound);
        }

        [Fact]
        public void Should_Authenticate_Exception()
        {
            var userViewModel = _userService.Authenticate(null);

            Assert.Null(userViewModel);
        }

        public void Dispose()
        {
            GC.SuppressFinalize(this);
        }
    }
}
