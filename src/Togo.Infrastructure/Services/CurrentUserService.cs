using System.Security.Claims;
using Microsoft.AspNetCore.Http;
using Microsoft.IdentityModel.JsonWebTokens;
using Togo.Core;
using Togo.Core.Interfaces;

namespace Togo.Infrastructure.Services;

public class CurrentUserService : ICurrentUserService
{
    private readonly IHttpContextAccessor _httpContextAccessor;

    public CurrentUserService(IHttpContextAccessor httpContextAccessor)
    {
        _httpContextAccessor = httpContextAccessor;
    }

    public string UserId
    {
        get
        {
            var userIdClaim = _httpContextAccessor?.HttpContext?.User.Claims.FirstOrDefault(claim =>
                claim.Type == ClaimTypes.NameIdentifier);

            return userIdClaim?.Value ?? string.Empty;
        }
    }

    public string SessionId
    {
        get
        {
            var sessionIdClaim = _httpContextAccessor?.HttpContext?.User.Claims.FirstOrDefault(claim =>
                claim.Type == JwtRegisteredClaimNames.Sid);

            return sessionIdClaim?.Value ?? string.Empty;
        }
    }
    
    public IList<string> RoleIds
    {
        get
        {
            var roleIdClaims =
                _httpContextAccessor?.HttpContext?.User.Claims.Where(claim => claim.Type == ClaimTypes.Role);
            var roleIds = roleIdClaims
                ?.Select(claim => claim.Value)
                .ToList();

            return roleIds ?? new List<string>();
        }
    }

    public int MaxTasksPerDay
    {
        get
        {
            var maxTasksPerDayClaim = _httpContextAccessor?.HttpContext?.User.Claims.FirstOrDefault(claim =>
                claim.Type == TogoCustomClaims.MaxTasksPerDay);
            
            return int.TryParse(maxTasksPerDayClaim?.Value, out var maxTasksPerDay) ? maxTasksPerDay : 0;
        }
    }
}
