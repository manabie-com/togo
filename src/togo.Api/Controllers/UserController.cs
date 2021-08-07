using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
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
        public async Task<ApiResponse<string>> Login([FromQuery] string user_id, [FromQuery] string password)
        {
            return await _userService.Login(user_id, password);
        }
    }
}
