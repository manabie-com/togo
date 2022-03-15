using Microsoft.AspNetCore.Mvc;
using Todo.Contract;
using Todo.Domain.Interfaces;

namespace Todo.Host.Controllers;

[ApiController]
[Route("api/todos")]
public class TodoAppController : ControllerBase
{
    private readonly ILogger<TodoAppController> _logger;

    private readonly IUserManagementService _userService;

    private readonly ITodoManagementService _todoService;

    public TodoAppController(ILogger<TodoAppController> logger, IUserManagementService userService, ITodoManagementService todoService)
    {
        _logger = logger ?? throw new ArgumentNullException(nameof(logger));
        _userService = userService ?? throw new ArgumentNullException(nameof(userService));
        _todoService = todoService ?? throw new ArgumentNullException(nameof(todoService));
    }

    [HttpGet]
    public IActionResult Get([FromQuery] long id)
    {
        try
        {
            if (id < 1)
            {
                return BadRequest();
            }
            
            return Ok();
        }
        catch (Exception e)
        {
            _logger.LogError(e, e.Message);
            return StatusCode(500, e.Message);
        }
    }

    [HttpDelete]
    public async Task<IActionResult> DeleteTodo([FromQuery] long id)
    {
        try
        {
            return Ok(await _todoService.DeleteAsync(id));
        }
        catch (Exception e)
        {
            _logger.LogError(e, e.Message);
            return StatusCode(500, e.Message);
        }
    }

    [HttpPost]
    public async Task<IActionResult> CreateTodo([FromBody] CreateTodoResource resource)
    {
        try
        {
            TryValidateModel(resource);

            if (!ModelState.IsValid)
            {
                return BadRequest();
            }

            var user = await _userService.GetAsync(resource.UserId);

            var tasksLimit = user.DailyTaskLimit;

            if (user.Todos != null)
            {
                var tasksToday = user.Todos.Count(t => t.DateCreatedUTC.Date == DateTime.UtcNow.Date);

                if (tasksToday >= tasksLimit)
                {
                    return BadRequest($"Unable to create a todo item because the user has exceeded the daily limit ({tasksLimit}).");
                }
            }

            var result = await _todoService.AddAsync(resource);

            var baseUri = "https://localhost:5001";
            var uri = $"{baseUri}/api/todos/{result.Id}";

            return Created(uri, result);
        }
        catch (Exception e)
        {
            _logger.LogError(e, e.Message);
            return StatusCode(500, e.Message);
        }
    }
}
