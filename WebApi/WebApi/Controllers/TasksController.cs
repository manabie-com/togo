using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System;
using System.Threading.Tasks;
using WebApi.Requests;
using WebApi.Services;

namespace WebApi.Controllers
{
    [ApiController]
    [Route("[controller]")]
    [Authorize]
    public class TasksController : ControllerBase
    {
        private readonly ITaskService _taskService;
        public TasksController(ITaskService taskService)
        {
            _taskService = taskService;
        }

        /// <summary>
        /// Create task API
        /// </summary>
        /// <param name="createTaskRequest"></param>
        /// <returns></returns>
        [HttpPost]
        public async Task<IActionResult> CreateTask([FromBody] CreateTaskRequest createTaskRequest)
        {
            var currentUserId = HttpContext.User.Identity.Name;
            if (string.IsNullOrEmpty(currentUserId))
            {
                return Unauthorized();
            }

            // Check limit daily task
            var isLimitTaskForToday = await _taskService.CheckLimitDailyTask(Guid.Parse(currentUserId));
            if (isLimitTaskForToday == null)
            {
                return BadRequest("User doesn't exist");
            }
            else if(isLimitTaskForToday == false)
            {
                return BadRequest("Daily task limit is reached");
            }

            var result = await _taskService.CreateTask(createTaskRequest, Guid.Parse(currentUserId));

            return Ok(result);
        }
    }
}
