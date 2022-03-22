using Common;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Models;
using Services.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using WebAPI.Extensions;
using WebAPI.Security;

namespace WebAPI.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class TasksController : ControllerBase
    {
        private readonly ITaskService _taskService;

        public TasksController(ITaskService taskService)
        {
            _taskService = taskService;
        }

        [HttpPost("create")]
        [Authorize]
        public int Create([FromBody] Tasks task)
        {
            return _taskService.Create(task, HttpContext.GetUserId());
        }

        [HttpGet("tasks-by-user-id")]
        [Authorize]
        public List<Tasks> GetTasksByUserId()
        {
            return _taskService.GetTasksByUserId(HttpContext.GetUserId());
        }
    }
}
