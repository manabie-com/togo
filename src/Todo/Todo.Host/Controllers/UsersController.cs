
using Microsoft.AspNetCore.Mvc;
using Todo.Contract;
using Todo.Domain.Interfaces;

namespace Todo.Host.Controllers;

[ApiController]
[Route("api/users")]
public class UsersController : ControllerBase
{
    private readonly ILogger<TodoAppController> _logger;

    private readonly IUserManagementService _service;

    public UsersController(ILogger<TodoAppController> logger, IUserManagementService service)
    {
        _logger = logger ?? throw new ArgumentNullException(nameof(logger));
        _service = service ?? throw new ArgumentNullException(nameof(service));
    }

    [HttpGet]
    public async Task<IActionResult> GetUser([FromQuery] long id)
    {
        try
        {
            if (id < 1)
            {
                return BadRequest();
            }
            
            return Ok(await _service.GetAsync(id));
        }
        catch (Exception e)
        {
            _logger.LogError(e, e.Message);
            return StatusCode(500, e.Message);
        }
    }

    [HttpDelete]
    public async Task<IActionResult> DeleteUser([FromQuery] long id)
    {
        try
        {
            if (id < 1)
            {
                return BadRequest();
            }

            return Ok(await _service.DeleteAsync(id));
        }
        catch (Exception e)
        {
            _logger.LogError(e, e.Message);
            return StatusCode(500, e.Message);
        }
    }

    [HttpPost]
    public async Task<IActionResult> CreateUser([FromBody] CreateUserResource resource)
    {
        try
        {
            TryValidateModel(resource);

            if (!ModelState.IsValid)
            {
                return BadRequest();
            }

            var result = await _service.AddAsync(resource);

            var baseUri = "https://localhost:5001";
            var uri = $"{baseUri}/api/users/{result}";

            return Created(uri, result);
        }
        catch (Exception e)
        {
            _logger.LogError(e, e.Message);
            return StatusCode(500, e.Message);
        }
    }
}
