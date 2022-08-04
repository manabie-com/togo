using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;
using Togo.Core.Interfaces;
using Togo.Infrastructure.Identities;
using Togo.Infrastructure.Persistence;

namespace Togo.Api.IntegrationTests;

public class TogoWebApplicationFactory : WebApplicationFactory<IntegrationTestStartup>
{
    
    
    protected override void ConfigureWebHost(IWebHostBuilder builder)
    {
        builder.ConfigureServices(services =>
        {
            services.AddControllers();

            services.AddDbContext<AppDbContext>(options =>
                options.UseNpgsql(
                    "User ID=postgres;Password=NoPassword1;Host=localhost;Port=5432;Database=togo.test;"));

            services.AddScoped<IAppDbContext, AppDbContext>();

            services.AddIdentityCore<AppUser>()
                .AddEntityFrameworkStores<AppDbContext>();

            services.AddScoped<IUserService, UserService>();
        });
    }
}