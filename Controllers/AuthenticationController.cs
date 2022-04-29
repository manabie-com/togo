using ManabieTodo.Models;
using ManabieTodo.Services;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;

namespace ManabieTodo.Controllers
{
    [Authorize]
    public class AuthenticationController : BaseController
    {
        private IAuthenticationService _authService { get; }

        public AuthenticationController(IAuthenticationService authService)
        {
            _authService = authService;
        }

        [AllowAnonymous]
        [HttpPost("login")]
        public IActionResult Login(UserModel model)
        {
            /** Imagination that I had encrypted the password string and now I had decrypted by some fabulous fucking algorithm */
            string decryptedPassword = model.Password;

            string jwt = _authService.Login(model.Username, decryptedPassword);

            if (string.IsNullOrEmpty(jwt))
            {
                return new ConflictObjectResult(new
                {
                    Message = "Đăng nhập thất bại"
                });
            }

            return new OkObjectResult(new
            {
                Token = jwt,
                Message = "Đăng nhập thành công"
            });
        }

        [HttpPost("fake-login")]
        [AllowAnonymous]
        public IActionResult FakeLogin([FromBody] string name)
        {
            string jwt = _authService.FakeLogin(name);

            if (string.IsNullOrEmpty(jwt))
            {
                return new ConflictObjectResult(new
                {
                    Message = "Đăng nhập fake thất bại"
                });
            }

            return new OkObjectResult(new
            {
                Token = jwt,
                Message = "Đăng nhập fake thành công"
            });
        }

        [HttpGet("logout")]
        public IActionResult Logout([FromHeader()] string authorization)
        {
            authorization = authorization.TrimStart(JwtBearerDefaults.AuthenticationScheme.ToArray<char>());
            authorization = authorization.Trim(' ');

            if (_authService.Logout(authorization))
            {
                return new OkObjectResult(new
                {
                    Message = "Đăng xuất thành công"
                });
            }

            return new ConflictObjectResult(new
            {
                Message = "Đăng xuất quài vậy ông nội"
            });
        }
    }
}