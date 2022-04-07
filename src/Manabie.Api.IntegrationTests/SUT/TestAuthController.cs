using System;
using System.Net;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Net.Mime;
using System.Security.Claims;
using System.Text.Encodings.Web;
using System.Text.Json.Serialization;
using System.Threading.Tasks;

using Manabie.Api.Insfrastructures;
using Manabie.Api.Middleware;
using Manabie.Api.Models;
using Manabie.Api.Services;
using Manabie.Api.Utilities;

using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.AspNetCore.TestHost;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using Microsoft.Net.Http.Headers;

using Newtonsoft.Json;

using Xunit;

namespace Manabie.Api.IntegrationTests.SUT;

public class TestAuthController
{

    [Fact]
    public async Task Get_SecurePageIsReturnedForAnAuthenticatedUser()
    {
        var factoty = new WebApplicationFactory<Program>();

        // Arrange
        var client = factoty.WithWebHostBuilder(builder =>
        {
            builder.ConfigureTestServices(services =>
            {
                var configuration = new ConfigurationBuilder()
                .SetBasePath(Environment.CurrentDirectory)
                .AddJsonFile("appsettings.json")
                .Build();

                services.AddAuthentication("Test")
                    .AddScheme<AuthenticationSchemeOptions, TestAuthHandler>(
                        "Test", options => { });
                services.AddDbContext<ManaDbContext>(options => options.UseInMemoryDatabase("integration_test_db"));
                services.Configure<AppSettings>(configuration.GetSection("AppSettings"));
                services.AddScoped<IUserService, UserService>();
                services.AddScoped<ITaskService, TaskService>();
                services.AddAuthentication();
            })
            .Configure(app =>
            {
                app.UseRouting();
                app.UseAuthentication();
                app.UseEndpoints(endpoints =>
                {
                    endpoints.MapControllers();
                });
            })
            .UseTestServer()
            .UseContentRoot(Environment.CurrentDirectory)
            .UseConfiguration(new ConfigurationBuilder()
            .SetBasePath(Environment.CurrentDirectory)
                .AddJsonFile("appsettings.json")
                .Build());


        })
            .CreateClient(new WebApplicationFactoryClientOptions
            {
                AllowAutoRedirect = false,
            });

        client.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Test");

        //Act
        var response = await client.GetAsync("api/Auth");
        response.EnsureSuccessStatusCode();

        // Assert
        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
    }

    public class TestAuthHandler : AuthenticationHandler<AuthenticationSchemeOptions>
    {
        public TestAuthHandler(IOptionsMonitor<AuthenticationSchemeOptions> options,
            ILoggerFactory logger, UrlEncoder encoder, ISystemClock clock)
            : base(options, logger, encoder, clock)
        {
        }

        protected override Task<AuthenticateResult> HandleAuthenticateAsync()
        {
            var claims = new[] { new Claim(ClaimTypes.Name, "Test User"), new Claim("Id", "1") };
            var identity = new ClaimsIdentity(claims, "Test");
            var principal = new ClaimsPrincipal(identity);
            var ticket = new AuthenticationTicket(principal, "Test");

            var result = AuthenticateResult.Success(ticket);
            var user = new Entities.User
            {
                Id = 1,
                Username = "Test User"
            };
            Context.Items["User"] = user;

            Utility.GenerateJwtToken(user, "This is my custom Secret key for authentication.");

            return Task.FromResult(result);
        }
    }

}
