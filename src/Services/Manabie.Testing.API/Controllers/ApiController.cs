using MediatR;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using System.Security.Claims;
using static OpenIddict.Abstractions.OpenIddictConstants;

namespace Manabie.Testing.API.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public abstract class ApiController : ControllerBase
    {
        private IMediator _mediator;
        protected IMediator Mediator => _mediator ??= HttpContext.RequestServices.GetService<IMediator>();

        protected string Role => HttpContext.User.FindFirst(Claims.Role).Value;
        protected string UserId => HttpContext.User.FindFirst(Claims.Subject).Value;
    }
}
