using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using TODO.Repositories.Interfaces;
using TODO.Repositories.Models.RequestModels;

namespace TODO.Api.Controllers
{
    [Route("api/todos")]
    [ApiController]
    public class TodosController : ControllerBase
    {
        private readonly ILogger<TodosController> _logger;
        private readonly ITodoRepository _todoRepo;
        private readonly IUserRepository _userRepo;

        public TodosController(ILogger<TodosController> logger, ITodoRepository todoRepo, IUserRepository userRepo)
        {
            _logger = logger ?? throw new ArgumentException(nameof(logger));
            _todoRepo = todoRepo ?? throw new ArgumentException(nameof(todoRepo));
            _userRepo = userRepo ?? throw new ArgumentException(nameof(userRepo));
        }

        [HttpPost]
        public async Task<IActionResult> CreateTodo([FromBody] CreateTodoRequest request)
        {
            try
            {
                var user = (await _userRepo.GetUsers(request.UserId)).SingleOrDefault();
                var taskLimit = user.UserTodoConfig.DailyTaskLimit;
                var taskCount = user.Todos.Count(t => t.DateCreated.Value.Date == DateTime.UtcNow.Date);

                if (taskCount >= taskLimit)
                    return BadRequest("Unable to creat TODO: User has exceeded daily TODOs.");

                var result = await _todoRepo.CreateTodo(request);

                var url = $"{Request.Protocol}://{Request.Host}/api/todos/{result.TodoId}";
                return Created(url, result);
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }
    }
}
