using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Togo.Core.Exceptions;
using Togo.Infrastructure.Identities;
using Togo.Infrastructure.Identities.Dtos;
using Togo.Infrastructure.Identities.Exceptions;

namespace Togo.Api.Controllers;

[Route("api/user")]
public class UserController : ApiController
{
    private readonly IUserService _userService;

    public UserController(IUserService userService)
    {
        _userService = userService;
    }

    [HttpPost]
    [Authorize(Roles = Roles.Admin)]
    public async Task<IActionResult> CreateAsync([FromBody] CreateUserDto input)
    {
        try
        {
            return Ok(await _userService.CreateAsync(input));
        }
        catch (UserCreationFailedException ex)
        {
            return BadRequest(ex.Result);
        }
    }

    [HttpGet]
    [Authorize(Roles = Roles.Admin)]
    public async Task<IActionResult> GetAllAsync()
    {
        return Ok(await _userService.GetAllAsync());
    }

    [HttpPost("login")]
    [AllowAnonymous]
    public async Task<IActionResult> LoginAsync([FromBody] LoginDto input)
    {
        try
        {
            return Ok(await _userService.AuthenticateAsync(input));
        }
        catch (InvalidLoginException ex)
        {
            return BadRequest(ex.Message);
        }
    }
}
