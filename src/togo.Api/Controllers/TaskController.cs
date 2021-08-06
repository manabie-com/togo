using Microsoft.AspNetCore.Mvc;
using togo.Service.Dto;
using togo.Service.Interface;
using System.Threading.Tasks;
using TaskEntity = togo.Service.Context.Task;
using System.Collections.Generic;

namespace togo.Api.Controllers
{
    [Route("tasks")]
    public class TaskController : ControllerBase
    {
        private readonly ITaskService _taskService;
        public TaskController(ITaskService taskService)
        {
            _taskService = taskService;
        }

        [HttpPost]
        public async Task<ApiResponse<TaskEntity>> CreateTask([FromBody] TaskCreateDto input)
        {
            return await _taskService.Create(input);
        }

        [HttpGet]
        public async Task<ApiResponse<List<TaskEntity>>> ListTasks([FromQuery] string created_date)
        {
            return await _taskService.List(created_date);
        }
    }
}
