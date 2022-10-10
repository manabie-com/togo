using Microsoft.AspNetCore.Mvc;
using Swashbuckle.AspNetCore.Annotations;
using Todo.Application.Dtos;
using Todo.Application.Services;

namespace Todo.Api.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class TasksController : ControllerBase
    {
        private readonly IUserTaskService _taskService;

        public TasksController(IUserTaskService taskService)
        {
            _taskService = taskService;
        }

        [HttpPost]
        [SwaggerResponse(200, Type = typeof(string))]
        public async Task<IActionResult> CreateTask([FromBody] CreateEditTaskDto dto)
        {
            var rs = await _taskService.CreateTaskAsync(dto);
            return Ok(rs);
        }
    }
}