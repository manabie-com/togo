using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.Collections.Generic;
using System.Threading.Tasks;
using togo.Service.Interface;

namespace togo.Api.Controllers
{
    public class UserController : ControllerBase
    {
        private readonly IUserService _userService;
        public UserController(IUserService userService)
        {
            _userService = userService;
        }

        [AllowAnonymous]
        [HttpGet("login")]
        public async Task<ActionResult<Dictionary<string, string>>> Login([FromQuery] string user_id, [FromQuery] string password)
        {
            var token = await _userService.Login(user_id, password);
            return new Dictionary<string, string> { { "data", token } };
        }
    }
}
