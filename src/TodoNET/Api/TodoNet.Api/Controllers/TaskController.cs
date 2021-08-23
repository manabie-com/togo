using Application.DTOs.Task;
using Application.Interfaces.Services;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.Threading.Tasks;

namespace TodoNet.Api.Controllers
{
    [Route("api/[controller]")]
    [Authorize]
    [ApiController]
    public class  TaskController : ControllerBase
    {
        private readonly ITaskServiceAsync _taskServiceAsync;

        public TaskController(ITaskServiceAsync taskServiceAsync)
        {
            _taskServiceAsync = taskServiceAsync;
        }

        [HttpPost]
        public async Task<IActionResult> Post(CreateTaskRequest request)
        {
            return Ok(await _taskServiceAsync.CreateTaskAsync(request));
        }
        [HttpGet("getAll")]
        public async Task<IActionResult> Get()
        {
            return Ok(await _taskServiceAsync.GetTasksAsync());
        }
    }
}
