
using System.IdentityModel.Tokens.Jwt;
using System.Text;

using Manabie.Api.Services;
using Manabie.Api.Utilities;

using Microsoft.Extensions.Options;
using Microsoft.IdentityModel.Tokens;

namespace Manabie.Api.Middleware;

public class JWTMiddleware
{
    private readonly RequestDelegate _next;
    private readonly AppSettings _appSettings;

    public JWTMiddleware(RequestDelegate next, IOptions<AppSettings> appSettings)
    {
        _next = next;
        _appSettings = appSettings.Value;
    }

    public async Task Invoke(HttpContext context, IUserService userService)
    {
        var token = context.Request.Headers["Authorization"].FirstOrDefault()?.Split(" ").Last();

        if (token != null)
            attachUserToContext(context, userService, token);

        await _next(context);
    }

    private void attachUserToContext(HttpContext context, IUserService userService, string token)
    {
        try
        {
            var tokenHandler = new JwtSecurityTokenHandler();
            var key = Encoding.UTF8.GetBytes(_appSettings.Secret);
            tokenHandler.ValidateToken(token, new TokenValidationParameters
            {
                ValidateIssuerSigningKey = true,
                IssuerSigningKey = new SymmetricSecurityKey(key),
                ValidateIssuer = false,
                ValidateAudience = false,
                ClockSkew = TimeSpan.Zero
            }, out SecurityToken validatedToken);

            var jwtToken = (JwtSecurityToken)validatedToken;
            var userName = int.Parse(jwtToken.Claims.First(x => x.Type == "Id").Value);

            // attach user to context on successful jwt validation
            context.Items["User"] = userService.GetById(userName);
        }
        catch
        {
            ///
            /// do nothing if jwt validation fails
            /// user is not attached to context so request won't have access to secure routes
            ///
            throw new UnauthorizedAccessException();
        }
    }
}
