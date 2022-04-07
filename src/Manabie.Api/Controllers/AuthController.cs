
using Manabie.Api.Auth;
using Manabie.Api.Models;
using Manabie.Api.Services;

using Microsoft.AspNetCore.Mvc;

namespace Manabie.Api.Controllers;

[ApiController]
[Route("api/[controller]")]
public class AuthController : ControllerBase
{
    private IUserService _userService;

    public AuthController(IUserService userService)
    {
        _userService = userService;
    }

    [HttpGet("/Index")]
    public IActionResult Index()
    {
        return Ok();
    }


    [HttpPost("authenticate")]
    [Consumes("application/json")]
    public IActionResult Authenticate(AuthenticateRequest model)
    {
        var response = _userService.Authenticate(model);

        if (response == null)
            return BadRequest(new { message = "Username or password is incorrect" });

        return Ok(response);
    }

    [Authorize]
    [HttpGet]
    public IActionResult GetAll()
    {
        var users = _userService.GetAll();
        return Ok(users);
    }
}
