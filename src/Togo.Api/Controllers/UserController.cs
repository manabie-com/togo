using Microsoft.AspNetCore.Mvc;
using Togo.Infrastructure.Identities;
using Togo.Infrastructure.Identities.Dtos;

namespace Togo.Api.Controllers;

[Route("user")]
public class UserController : ApiController
{
    private readonly IUserService _userService;

    public UserController(IUserService userService)
    {
        _userService = userService;
    }

    [HttpPost]
    public async Task<IActionResult> CreateAsync([FromBody] CreateUserDto input)
    {
        return Ok(await _userService.CreateAsync(input));
    }

    [HttpGet]
    public async Task<IActionResult> GetAllAsync()
    {
        return Ok(await _userService.GetAllAsync());
    }

    [HttpPost("login")]
    public async Task<IActionResult> LoginAsync([FromBody] LoginDto input)
    {
        if (await _userService.AuthenticateAsync(input))
        {
            return NoContent();
        }

        return BadRequest();
    }
}