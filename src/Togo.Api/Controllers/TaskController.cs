using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Togo.Core.AppServices.TaskItems.Dtos;
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
        return Ok(await _taskItemAppService.CreateAsync(input));
    }
}
