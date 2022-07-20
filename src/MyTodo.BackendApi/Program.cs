using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;
using MyTodo.Data.EntityFramework;
using Microsoft.EntityFrameworkCore;

namespace MyTodo.BackendApi
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var host = CreateHostBuilder(args).Build();

            using (var scope = host.Services.CreateScope())
            {

                var services = scope.ServiceProvider;
                try
                {
                    var db = services.GetRequiredService<MyTodoDbContext>();
                    db.Database.Migrate();
                    //Seed
                    try
                    {
                        var initializer = services.GetService<DatabaseInitializer>();
                        initializer.Seed().Wait();
                    }
                    catch (Exception ex)
                    {
                        var logger = services.GetService<ILogger<Program>>();
                        logger.LogError(ex, "An error occurred while seeding the database");
                    }
                }
                catch (Exception ex)
                {
                    var logger = services.GetRequiredService<ILogger<Program>>();
                    logger.LogError(ex, "An error occurred while migrating the database.");
                }
                 
            }

            host.Run();
        }

        public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureWebHostDefaults(webBuilder =>
                {
                    webBuilder.UseStartup<Startup>();
                });
    }
}
