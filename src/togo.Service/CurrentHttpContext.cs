using Microsoft.AspNetCore.Http;
using System.Linq;
using System.Security.Claims;
using togo.Service.Interface;

namespace togo.Service
{
    public class CurrentHttpContext : ICurrentHttpContext
    {
        private readonly IHttpContextAccessor _httpContextAccessor;
        public CurrentHttpContext(IHttpContextAccessor httpContextAccessor)
        {
            _httpContextAccessor = httpContextAccessor;
        }

        public string GetCurrentUserId()
        {
            var userId = _httpContextAccessor.HttpContext.User?.Claims?
                                .FirstOrDefault(x => x.Type == ClaimTypes.NameIdentifier)?.Value;

            return userId;
        }
    }
}
