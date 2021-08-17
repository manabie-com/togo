using Microsoft.AspNetCore.Mvc;
using System;
using System.Threading.Tasks;
using WebApi.Services;

namespace WebApi.Controllers
{
    [ApiController]
    public class UserController : ControllerBase
    {
        private readonly IUserService _userService;
        public UserController(IUserService userService)
        {
            _userService = userService;
        }

        [HttpGet("login")]
        public async Task<IActionResult> Login(Guid user_id, string password)
        {
            if (user_id == Guid.Empty || String.IsNullOrEmpty(password))
            {
                return BadRequest("User Id or password is empty");
            }
            var result = await _userService.Login(user_id, password);

            if (result == null)
            {
                return BadRequest("User doesn't exist");
            }

            return Ok(result);
        }
    }
}
