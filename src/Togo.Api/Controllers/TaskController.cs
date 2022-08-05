using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Togo.Core.AppServices.TaskItems.Dtos;
using Togo.Core.Exceptions;
using Togo.Core.Interfaces.AppServices;

namespace Togo.Api.Controllers;

[Route("api/tasks")]
public class TaskController : ApiController
{
    private readonly ITaskItemAppService _taskItemAppService;

    public TaskController(ITaskItemAppService taskItemAppService)
    {
        _taskItemAppService = taskItemAppService;
    }

    [HttpPost]
    [Authorize]
    public async Task<IActionResult> CreateAsync(CreateTaskItemDto input)
    {
        try
        {
            return Ok(await _taskItemAppService.CreateAsync(input));
        }
        catch (TaskLimitExceededException ex)
        {
            return BadRequest($"Task limit of {ex.TaskLimit} exceeded");
        }
    }
}
