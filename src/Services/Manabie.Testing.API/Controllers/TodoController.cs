using Manabie.Testing.Application.Common.Models;
using Manabie.Testing.Application.Todos.Commands.CreateTodoItem;
using Manabie.Testing.Application.Todos.Queries.GetAllTodos;
using Manabie.Testing.Application.UserLimits.Commands.CreateUserLimit;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using OpenIddict.Validation.AspNetCore;

namespace Manabie.Testing.API.Controllers
{
    [ApiController]
    [Route("[controller]/[action]")]
    [Authorize(AuthenticationSchemes = OpenIddictValidationAspNetCoreDefaults.AuthenticationScheme)]
    public class TodoController : ApiController
    {
        [HttpPost]
        public async Task<IActionResult> AddTodoAsync(CreateTodoItemCommand command)
        {
            if(command == null) return BadRequest();

            command.UserId = UserId;
            command.Role = Role;

            var userLimitId = await Mediator.Send(new CreateUserLimitCommand()
            {
                UserId = command.UserId,
                Role = command.Role,
            });

            var result = await Mediator.Send(command);

            return Ok(result);
        }

        [HttpGet]
        public async Task<IActionResult> GetTodosAsync()
        {
            var todos = await Mediator.Send(new GetAllTodoQuery
            {
                UserId = UserId
            });

            return Ok(todos);
        }
    }
}
