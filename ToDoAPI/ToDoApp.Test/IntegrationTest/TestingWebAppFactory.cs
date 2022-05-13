using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.VisualStudio.TestPlatform.TestHost;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using ToDoApp.API;
using ToDoApp.DTO;

namespace ToDoApp.Test.IntegrationTest
{
    public class TestingWebAppFactory<TEntryPoint> : WebApplicationFactory<Startup> where TEntryPoint : Startup
    {
        protected override void ConfigureWebHost(IWebHostBuilder builder)
        {
            builder.ConfigureServices(services =>
            {
                var descriptior = services.SingleOrDefault(d =>
                d.ServiceType == typeof(DbContextOptions<DatabaseContext>));

                if(descriptior != null)
                    services.Remove(descriptior);

                services.AddDbContext<DatabaseContext>(options =>
                {
                    options.UseInMemoryDatabase(Guid.NewGuid().ToString());
                });

                
                var serviceProvider = services.BuildServiceProvider();
                using(var scope = serviceProvider.CreateScope())
                using(var context = serviceProvider.GetRequiredService<DatabaseContext>())
                {
                    try
                    {
                        context.Database.EnsureCreated();
                        context.Database.EnsureDeleted();
                        context.CreateMockSampleData();
                    }
                    catch(Exception ex)
                    {
                        throw;
                    }
                }
            });
        }
    }
}
