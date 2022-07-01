using Manabie.BasicIdentityServer.Application.Common.Interfaces;
using Manabie.BasicIdentityServer.Infrastructure.Identity;
using Manabie.BasicIdentityServer.Infrastructure.Persistence;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.IdentityModel.Tokens;
using OpenIddict.Abstractions;
using Quartz;
using static OpenIddict.Abstractions.OpenIddictConstants;

namespace Manabie.BasicIdentityServer.Infrastructure
{
    public static class ConfigurationServices
    {
        public static IServiceCollection AddInfrastructure(this IServiceCollection services)
        {
            //services.AddScoped<AuditableEntitySaveChangesInterceptor>();

            services.AddDbContext<ApplicationDbContext>(options =>
            {
                options.UseInMemoryDatabase("CleanArchitectureDb");

                // Register the entity sets needed by OpenIddict.
                // Note: use the generic overload if you need
                // to replace the default OpenIddict entities.
                options.UseOpenIddict();

            });

            services
                .AddIdentity<ApplicationUser, IdentityRole>()
                .AddEntityFrameworkStores<ApplicationDbContext>()
                .AddDefaultTokenProviders();

            services.AddQuartz(options =>
            {
                options.UseMicrosoftDependencyInjectionJobFactory();
                options.UseSimpleTypeLoader();
                options.UseInMemoryStore();
            });

            services.AddBasicAuthenticationServer();

            services.AddScoped<IApplicationDbContext>(provider => provider.GetRequiredService<ApplicationDbContext>());

            services.AddTransient<IIdentityService, IdentityService>();

            services.AddScoped<DataSeed>();

            return services;
        }

        internal static IServiceCollection AddBasicAuthenticationServer(this IServiceCollection services)
        {

            // Configure Identity to use the same JWT claims as OpenIddict instead
            // of the legacy WS-Federation claims it uses by default (ClaimTypes),
            // which saves you from doing the mapping in your authorization controller.
            services.Configure<IdentityOptions>(options =>
            {
                options.ClaimsIdentity.UserNameClaimType = Claims.Name;
                options.ClaimsIdentity.UserIdClaimType = Claims.Subject;
                options.ClaimsIdentity.RoleClaimType = Claims.Role;
            });

            services.AddOpenIddict()
             // Register the OpenIddict core components.
             .AddCore(options =>
             {
                 // Configure OpenIddict to use the Entity Framework Core stores and models.
                 // Note: call ReplaceDefaultEntities() to replace the default entities.
                 options.UseEntityFrameworkCore()
                     .UseDbContext<ApplicationDbContext>();

                 options.UseQuartz();

             })
             // Register the OpenIddict server components.
             .AddServer(options =>
             {
                 // Enable the token endpoint.
                 options.SetTokenEndpointUris("/api/connect/token");
                 // Enable the client credentials flow.
                 options.AllowPasswordFlow();

                 options.AddEncryptionKey(new SymmetricSecurityKey(
                                Convert.FromBase64String("DRjd/GnduI3Efzen9V9BvbNUfc/VKgXltV7Kbk9sMkY=")));

                 options.DisableAccessTokenEncryption();

                 //options.RegisterScopes(OpenIddictConstants.Scopes.Email,
                 //                       OpenIddictConstants.Scopes.Profile,
                 //                       OpenIddictConstants.Scopes.Roles);

                 // Register the signing and encryption credentials.
                 options
                        .AddDevelopmentSigningCertificate();

                 // Accept anonymous clients (i.e clients that don't send a client_id).
                 options.AcceptAnonymousClients();

                 // Register the ASP.NET Core host and configure the ASP.NET Core-specific options.
                 options.UseAspNetCore()
                        .EnableTokenEndpointPassthrough();
             })
             .AddValidation(options =>
             {
                 //Import the configuration from the local OpenIddict server instance.
                 options.UseLocalServer();

                 // Register the ASP.NET Core host.
                 options.UseAspNetCore();
             }); 

            ///cross domain
            services.AddCors(o => o.AddPolicy("MyPolicy", builder =>
            {
                builder.AllowAnyHeader().AllowAnyMethod().AllowAnyOrigin();
            }));

            //policy
            //services
            //    .AddAuthorization(options => options.AddPolicy("CanPurge", policy => policy.RequireRole("Administrator")));

            return services;
        }
    }
}