using Application.DTOs.User;
using Application.Interfaces.Repositories;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.Threading.Tasks;

namespace TodoNet.Api.Controllers
{
    [Route("api/[controller]")]
    [Authorize]
    [ApiController]
    public class UserController : ControllerBase
    {
        private readonly IUserServiceAsync _userServiceAsync;
        public UserController(IUserServiceAsync userServiceAsync)
        {
            _userServiceAsync = userServiceAsync;
        }

        [AllowAnonymous]
        [HttpPost("login")]
        public async Task<IActionResult> AuthenticateAsync(AuthenticationRequest request)
        {
            return Ok(await _userServiceAsync.AuthenticateAsync(request));
        }
        [HttpPost]
        public async Task<IActionResult> Post(CreateUserRequest request)
        {
            return Ok(await _userServiceAsync.CreateUserAsync(request));
        }
        [HttpGet("getAll")]
        public async Task<IActionResult> Get()
        {
            return Ok(await _userServiceAsync.GetUsers());
        }
    }
}
