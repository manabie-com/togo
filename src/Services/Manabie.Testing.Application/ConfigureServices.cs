using Microsoft.Extensions.DependencyInjection;
using MediatR;
using System.Reflection;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using OpenIddict.Validation.AspNetCore;
using Microsoft.IdentityModel.Tokens;

namespace Manabie.Testing.Application
{
    public static class ConfigureServices
    {
        public static IServiceCollection AddApplication(this IServiceCollection services)
        {
            services.AddMediatR(Assembly.GetExecutingAssembly());

            services
                .AddOpenIddict()
                .AddValidation(options =>
                {
                    // Note: the validation handler uses OpenID Connect discovery
                    // to retrieve the issuer signing keys used to validate tokens.
                    options.SetIssuer("https://localhost:7173/");
                    //options.AddAudiences("resource_server");

                    // Register the encryption credentials. This sample uses a symmetric
                    // encryption key that is shared between the server and the Api2 sample
                    // (that performs local token validation instead of using introspection).
                    //
                    // Note: in a real world application, this encryption key should be
                    // stored in a safe place (e.g in Azure KeyVault, stored as a secret).
                    options.AddEncryptionKey(new SymmetricSecurityKey(
                        Convert.FromBase64String("DRjd/GnduI3Efzen9V9BvbNUfc/VKgXltV7Kbk9sMkY=")));

                    // Register the System.Net.Http integration.
                    options.UseSystemNetHttp();
                    // Register the ASP.NET Core host.
                    options.UseAspNetCore();
                });

            services
                .AddAuthentication(OpenIddictValidationAspNetCoreDefaults.AuthenticationScheme);

            services.AddAuthorization();
            return services;
        }
    }
}