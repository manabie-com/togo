namespace ManabieTodo.Utils
{
    public class ValidTokenMiddleware
    {
        private readonly RequestDelegate _next;


        public ValidTokenMiddleware(RequestDelegate next)
        {
            _next = next;
        }

        public async Task InvokeAsync(HttpContext context)
        {
            string authQuery = context.Request.Headers["Authorization"];

            if (!string.IsNullOrWhiteSpace(authQuery))
            {

            }

            await _next(context);
        }
    }

    public static class ValidTokenMiddlewareExtensions
    {
        public static IApplicationBuilder UseValidTokenMiddleware(
            this IApplicationBuilder builder)
        {
            return builder.UseMiddleware<ValidTokenMiddleware>();
        }
    }
}