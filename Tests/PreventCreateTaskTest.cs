using System.Security.Claims;
using ManabieTodo.Constants;
using ManabieTodo.Utils;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Abstractions;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.AspNetCore.Mvc.ModelBinding;
using Xunit;

public class PreventTaskCreate
{
    [Fact]
    public void a()
    {
        //Arrange
        var modelState = new ModelStateDictionary();
        var httpContext = new DefaultHttpContext();
        var claimsPrincipal = new ClaimsPrincipal(
            new ClaimsIdentity(
                new[] {
                        new Claim(ClaimTag.Id, "1"),
                        new Claim(ClaimTag.Name,  "Unknown"),
                        new Claim(ClaimTag.AllowedTaskDay, "5"),
                    }
                )
            );

        httpContext.User = claimsPrincipal;

        var context = new ResourceExecutingContext(
            new ActionContext(
                httpContext: httpContext,
                routeData: new RouteData(),
                actionDescriptor: new ActionDescriptor(),
                modelState: modelState
            ),
            new List<IFilterMetadata>(),
            new List<IValueProviderFactory>());

        var preventTask = new PreventTaskCreateAttribute();

        //Act
        for (int i = 0; i <= 6; i++)
        {
            preventTask.OnResourceExecuting(context);
            var conflitResult = context.Result as ConflictObjectResult;

            if (i < 5)
            {
                Assert.True(conflitResult == null);
            }
            else
            {
                Assert.False(conflitResult == null);
            }
        }
    }
}