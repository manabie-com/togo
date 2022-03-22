using Microsoft.AspNetCore.Mvc;
using Services.Interfaces;
using Services.Models;
using Services.ViewModels;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using WebAPI.Security;

namespace WebAPI.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class UsersController : ControllerBase
    {
        private readonly IUserService _userService;
        private readonly IJwtHandler _jwtHandler;

        public UsersController(IUserService userService,
                               IJwtHandler jwtHandler)
        {
            _userService = userService;
            _jwtHandler = jwtHandler;
        }

        [HttpPost("login")]
        public string Login(UserLogin userLogin)
        {
            var user = _userService.Authenticate(userLogin);

            if (user != null) return _jwtHandler.Create(user);
            
            return "Unauthorized";
        }

        [HttpPost("register")]
        public int Register(UserRegister userRegister)
        {
            return _userService.Register(userRegister) ? 1 : 0;
        }
    }
}
