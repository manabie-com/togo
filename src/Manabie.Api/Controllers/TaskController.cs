using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Threading.Tasks;

using Manabie.Api.Entities;
using Manabie.Api.Models;
using Manabie.Api.Services;

using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace Manabie.Api.Controllers;

[Route("api/[controller]")]
[ApiController]
public class TaskController : BaseController
{
    private readonly ITaskService _taskService;

    public TaskController(ITaskService taskService)
    {
        _taskService = taskService;
    }

    [HttpGet]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(IEnumerable<Entities.Task>))]
    public async Task<ActionResult> Get()
    {
        return Ok(_taskService.GetTasks(UserInfo.Username));
    }
    [HttpPost]
    [ProducesResponseType(StatusCodes.Status200OK, Type = typeof(Entities.Task))]
    public async Task<ActionResult> Post(TaskViewModel task)
    {
        return Ok(_taskService.AddTask(UserInfo.Username, task));
    }
}
