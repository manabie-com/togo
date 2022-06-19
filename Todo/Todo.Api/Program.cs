//using Newtonsoft.Json;
//using Todo.Api;
//using Todo.Application;
//using Todo.Application.Extensions;
//using Todo.Infrastructure;

//var builder = WebApplication.CreateBuilder(args);
//WebApplication.User
//var Configuration = builder.Configuration;

//var b = builder.Services.BuildServiceProvider().GetService<Todo.Application.Services.IUserTaskService>();
//var dbSeeder = builder.Services.BuildServiceProvider().GetService<ApplicationDbSeeder>();
//var dbSeer = new ApplicationDbSeeder();
//var dbSeeder = builder.
//dbSeeder.EnsureMigrate();

//dbSeeder.EnsureData();

//var app = Startup.InitializeApp(args);
//app.Run();
using Todo.Api;

public class Program
{
    public static void Main(string[] args)
    {
        CreateHostBuilder(args).Build().Run();
    }

    public static IHostBuilder CreateHostBuilder(string[] args) =>
        Host.CreateDefaultBuilder(args)
            .ConfigureWebHostDefaults(webBuilder =>
            {
                webBuilder.UseStartup<Startup>();
            });
}