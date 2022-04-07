
using Manabie.Api.Entities;

using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;

namespace Manabie.Api.Auth;

public class AuthorizeAttribute : Attribute, IAuthorizationFilter
{
    public void OnAuthorization(AuthorizationFilterContext context)
    {
        var user = (User)context.HttpContext.Items["User"];
        if (user == null)
        {
            // not logged in
            context.Result = new JsonResult(new { message = "Unauthorized" }) { StatusCode = StatusCodes.Status401Unauthorized };
        }
    }
}
