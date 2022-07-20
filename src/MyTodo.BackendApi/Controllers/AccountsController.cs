using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Logging;
using Microsoft.IdentityModel.Tokens;
using MyTodo.BackendApi.Models;
using MyTodo.Data.Entities;
using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Linq;
using System.Security.Claims;
using System.Text;
using System.Threading.Tasks;

namespace MyTodo.BackendApi.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    [Authorize]
    public class AccountsController : ControllerBase
    {
        private readonly UserManager<AppUser> _userManager;
        private readonly RoleManager<AppRole> _roleManager;
        private readonly SignInManager<AppUser> _signInManager;
        private readonly IConfiguration _config;

        public AccountsController(
            UserManager<AppUser> userManager,
            RoleManager<AppRole> roleManager,
            SignInManager<AppUser> signInManager,
            IConfiguration config)
        {
            _userManager = userManager;
            _roleManager = roleManager;
            _signInManager = signInManager;
            _config = config;
        }
        [HttpGet]
        public async Task<ActionResult> GetAll()
        {
            var users = _userManager.Users.Select(x => new { x.Id, x.Email, x.UserName, x.TaskCount, x.TaskLimit });
            return Ok(users);
        }
        [HttpGet("{email}")]
        public async Task<IActionResult> GetById(string email)
        {
            var result = await _userManager.FindByEmailAsync(email);
            var rolesList = await _userManager.GetRolesAsync(result);

            var resultobject = new
            {
                Id = result.Id,
                Email = result.Email,
                UserName = result.UserName,
                TaskCount = result.TaskCount,
                TaskLimit = result.TaskLimit,
                Roles=string.Join(';', rolesList)
            };
            return Ok(resultobject);
        }
        [HttpPost]
        [AllowAnonymous]
        [Route("login")]

        public async Task<IActionResult> Login([FromBody] LoginViewModel model)
        {
            if (!ModelState.IsValid)
            {
                return new BadRequestObjectResult(model);
            }
            var user = await _userManager.FindByNameAsync(model.UserName);
            if (user != null)
            {
                var result = await _signInManager.PasswordSignInAsync(model.UserName, model.Password, false, true);
                if (!result.Succeeded)
                {
                    return new BadRequestObjectResult(result.ToString());
                }
                var roles = await _userManager.GetRolesAsync(user);
                var claims = new[]
                {
                    new Claim(JwtRegisteredClaimNames.Email, user.Email),
                    new Claim(JwtRegisteredClaimNames.UniqueName, user.UserName),
                    new Claim(ClaimTypes.NameIdentifier, user.Id.ToString()),
                    new Claim("roles", string.Join(";",roles)),
                    new Claim(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString())
                };
                var key = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_config["Tokens:Key"]));
                var creds = new SigningCredentials(key, SecurityAlgorithms.HmacSha256);

                var token = new JwtSecurityToken(_config["Tokens:Issuer"],
                    _config["Tokens:Issuer"],
                    claims,
                    expires: DateTime.UtcNow.AddMinutes(60),
                    signingCredentials: creds);

                return new OkObjectResult(new
                {
                    token = new JwtSecurityTokenHandler().WriteToken(token),
                    expiration = token.ValidTo
                });

            }
            return new BadRequestObjectResult("Login failure");
        }

        [HttpPost]
        [AllowAnonymous]
        [Route("register")]
        public async Task<IActionResult> Register([FromBody] RegisterViewModel model)
        {
            var userexist = await _userManager.FindByNameAsync(model.Email);
            if (userexist != null)
            {
                return StatusCode(StatusCodes.Status500InternalServerError, new ApiResponse()
                {
                    Status = "Error",
                    Message = "User already exists!"
                });
            }
            var user = new AppUser
            {
                UserName = model.Email,
                Email = model.Email,
                TaskCount = 0,
                TaskLimit = 5
            };
            var result = await _userManager.CreateAsync(user, model.Password);
            if (!result.Succeeded)
            {
                return StatusCode(StatusCodes.Status500InternalServerError, new ApiResponse()
                {
                    Status = "Error",
                    Message = "User creation failed! Please check user detail and try again."
                });

            }
            return Ok(new ApiResponse()
            {
                Status = "Success",
                Message = "User created successfully!"
            });
        }
        [HttpPut]
        [Route("update")]
        public async Task<IActionResult> Update([FromBody] UpdateUserViewModel model)
        {
            var user = await _userManager.FindByIdAsync(model.UserId.ToString());
            if (model.TaskLimit <= user.TaskLimit)
            {
                return StatusCode(StatusCodes.Status500InternalServerError, new ApiResponse()
                {
                    Status = "Error",
                    Message = "TaskLimit invalid."
                });
            }
            user.TaskLimit = model.TaskLimit;
            await _userManager.UpdateAsync(user);
            return Ok(new ApiResponse()
            {
                Status = "Success",
                Message = "User updated successfully!"
            });
        }

    }
}
