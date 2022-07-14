using MediatR;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace Manabie.BasicIdentityServer.API.Controllers
{
    [Authorize]
    [Route("api/[controller]")]
    [ApiController]
    public class TodosController : ControllerBase
    {
        //[HttpPost]
        //public async Task<ActionResult<int>> Create(CreateTodoItemCommand command)
        //{
        //    return await Mediator.Send(command);
        //}

    }
}
