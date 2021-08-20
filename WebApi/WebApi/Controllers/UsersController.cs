using Microsoft.AspNetCore.Mvc;
using System;
using System.Threading.Tasks;
using WebApi.Services;

namespace WebApi.Controllers
{
    [ApiController]
    public class UsersController : ControllerBase
    {
        private readonly IUserService _userService;
        public UsersController(IUserService userService)
        {
            _userService = userService;
        }

        /// <summary>
        /// Login API
        /// </summary>
        /// <param name="user_id"></param>
        /// <param name="password"></param>
        /// <returns></returns>
        [HttpGet("login")]
        public async Task<IActionResult> Login(Guid user_id, string password)
        {
            if (user_id == Guid.Empty || String.IsNullOrEmpty(password))
            {
                return BadRequest("User Id or Password is empty");
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
