using System.Reflection;
using System.Text;
using FluentValidation;
using FluentValidation.AspNetCore;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.IdentityModel.Tokens;
using Togo.Core.Interfaces;
using Togo.Infrastructure.Identities;
using Togo.Infrastructure.Persistence;
using Togo.Infrastructure.Services;

namespace Togo.Infrastructure;

public static class DependencyInjection
{
    public static IServiceCollection AddInfrastructureServices(this IServiceCollection services, TogoAppSettings togoAppSettings)
    {
        services.AddControllers().AddFluentValidation();
        services.AddValidatorsFromAssembly(Assembly.GetExecutingAssembly());
        services.AddEndpointsApiExplorer();
        services.AddSwaggerGen();
        services.AddDataAccess(togoAppSettings);
        services.AddIdentityFramework();
        services.AddAuthentication(togoAppSettings);
        services.AddHttpContextAccessor();
        services.AddScoped<IUserService, UserService>();
        services.AddSingleton<ICurrentUserService, CurrentUserService>();
        return services;
    }

    private static IServiceCollection AddAuthentication(this IServiceCollection services, TogoAppSettings togoAppSettings)
    {
        services
            .AddAuthentication(authenticationOptions =>
            {
                authenticationOptions.DefaultAuthenticateScheme = JwtBearerDefaults.AuthenticationScheme;
                authenticationOptions.DefaultChallengeScheme = JwtBearerDefaults.AuthenticationScheme;
            })
            .AddJwtBearer(jwtBearerOptions =>
            {
                jwtBearerOptions.RequireHttpsMetadata = false;
                jwtBearerOptions.SaveToken = true;
                jwtBearerOptions.TokenValidationParameters = new TokenValidationParameters
                {
                    ValidateIssuer = true,
                    ValidIssuer = togoAppSettings.JwtBearer.Issuer,
                    ValidateIssuerSigningKey = true,
                    IssuerSigningKey = new SymmetricSecurityKey(
                        Encoding.ASCII.GetBytes(
                            togoAppSettings.JwtBearer.SecurityKey)),
        
                    ValidateAudience = true,
                    ValidAudience = togoAppSettings.JwtBearer.Audience
                };
            });
        
        return services;
    }
    
    private static IServiceCollection AddDataAccess(this IServiceCollection services, TogoAppSettings togoAppSettings)
    {
        services.AddDbContext<AppDbContext>(options =>
            options.UseNpgsql(togoAppSettings.ConnectionStrings.Default));

        services.AddScoped<IUnitOfWork, EfUnitOfWork>();
        
        return services;
    }
    
    private static IServiceCollection AddIdentityFramework(this IServiceCollection services)
    {
        services.AddIdentityCore<AppUser>()
            .AddRoles<IdentityRole>()
            .AddRoleManager<RoleManager<IdentityRole>>()
            .AddEntityFrameworkStores<AppDbContext>();
        
        return services;
    }
}
