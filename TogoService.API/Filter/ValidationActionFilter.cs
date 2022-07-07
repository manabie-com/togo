using System.Collections.Generic;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using TogoService.API.Dto;
using TogoService.API.Infrastructure.Helper.MessageUtil;

namespace TogoService.API.Filter
{
    public class ValidationActionFilter : IActionFilter
    {
        public void OnActionExecuted(ActionExecutedContext context)
        {
            // Not doing anything after the action has executed
        }
        public void OnActionExecuting(ActionExecutingContext context)
        {
            var logger = context.HttpContext.RequestServices.GetService<ILogger<Microsoft.AspNetCore.Mvc.Controller>>();
            var modelState = context.ModelState;
            if (!modelState.IsValid)
            {
                var errors = new List<string>();
                foreach (var state in modelState)
                {
                    foreach (var error in state.Value.Errors)
                    {
                        errors.Add(error.ErrorMessage);
                    }
                }
                if (logger != null)
                {
                    logger.LogError(string.Join(", ", errors));
                }


                context.Result = new BadRequestObjectResult(
                    new CommonResponse<string>(
                        StatusCodes.Status400BadRequest,
                        CommonMessages.InvalidObject));
            }
        }
    }
}
