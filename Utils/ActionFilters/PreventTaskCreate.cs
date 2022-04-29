using ManabieTodo.Constants;
using ManabieTodo.Services;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;

namespace ManabieTodo.Utils
{
    public class PreventTaskCreateAttribute : Attribute, IResourceFilter, IAsyncResourceFilter
    {
        private IFakeCache _fakeCache { get; }

        public PreventTaskCreateAttribute()
        {
            _fakeCache = CacheFactory.FakeCache;
        }

        public void OnResourceExecuting(ResourceExecutingContext context)
        {
            HttpContext httpContext = context.HttpContext;

            int userId = int.Parse(httpContext.User.FindFirst(ClaimTag.Id).Value);
            int allowedTaskDay = int.Parse(httpContext.User.FindFirst(ClaimTag.AllowedTaskDay).Value);

            int cntTaskCreated = _fakeCache.Tasks.Where(t =>
                string.Compare(t.GetString("CREATE_DATE"), DateTime.Now.ToString("YYYY-MM-dd"), true) == 0 &&
                t.GetInt("USER_ID") == userId
            ).Count();

            if (cntTaskCreated >= allowedTaskDay)
            {
                context.Result = new ConflictObjectResult(
                    new
                    {
                        Message = "Siêng tạo task vậy bạn!?!"
                    });
            }

            _fakeCache.Task = new Dictionary<string, object>
            {
                ["USER_ID"] = userId,
            };

        }

        public void OnResourceExecuted(ResourceExecutedContext context)
        {
        }

        public async Task OnResourceExecutionAsync(ResourceExecutingContext context, ResourceExecutionDelegate next)
        {
            OnResourceExecuting(context);

            var conflitResult = context.Result as ConflictObjectResult;

            if (conflitResult != null)
            {
                return;
            }


            await next();
        }
    }
}