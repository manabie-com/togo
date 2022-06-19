using Newtonsoft.Json;
using Todo.Application;
using Todo.Application.Extensions;
using Todo.Infrastructure;

namespace Todo.Api
{
    //public static class Startup
    //{
    //    public static WebApplication InitializeApp(string[] args)
    //    {
    //        var builder = WebApplication.CreateBuilder(args);
    //        ConfigureServices(builder);

    //        var app = builder.Build();
    //        Configure(app);

    //        return app;
    //    }
    //    private static void ConfigureServices(WebApplicationBuilder builder)
    //    {
    //        var Configuration = builder.Configuration;

    //        // Add services to the container.

    //        builder.Services.AddControllers();
    //        // Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
    //        builder.Services.AddEndpointsApiExplorer();
    //        builder.Services.AddSwaggerGen();

    //        builder.Services.AddInfrastructure(Configuration);

    //        builder.Services.AddApplication(Configuration);

    //        builder.Services.AddControllers().AddNewtonsoftJson(options =>
    //        {
    //            options.SerializerSettings.ReferenceLoopHandling = ReferenceLoopHandling.Ignore;
    //        });

    //        ResolverFactory.ServiceCollection = builder.Services;
    //    }
    //    private static void Configure(WebApplication app)
    //    {

    //        // Configure the HTTP request pipeline.
    //        if (app.Environment.IsDevelopment())
    //        {
    //            app.UseSwagger();
    //            app.UseSwaggerUI();
    //        }

    //        app.UseHttpsRedirection();

    //        app.UseAuthorization();

    //        app.MapControllers();
    //    }
    //}


    public class Startup
    {
        public IConfiguration Configuration { get; }
        public IWebHostEnvironment Environment { get; }

        public Startup(IConfiguration configuration, IWebHostEnvironment environment)
        {
            Configuration = configuration;
            Environment = environment;
        }

        public void ConfigureServices(IServiceCollection services)
        {
            services.AddControllers();
            // Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
            services.AddEndpointsApiExplorer();
            services.AddSwaggerGen();

            services.AddInfrastructure(Configuration);

            services.AddApplication(Configuration);

            services.AddControllers().AddNewtonsoftJson(options =>
            {
                options.SerializerSettings.ReferenceLoopHandling = ReferenceLoopHandling.Ignore;
            });

            ResolverFactory.ServiceCollection = services;
        }

        public void Configure(IApplicationBuilder app, IWebHostEnvironment env, ApplicationDbSeeder dbSeeder)
        {
            dbSeeder.EnsureMigrate();

            dbSeeder.EnsureData();

            // Configure the HTTP request pipeline.
            if (env.IsDevelopment())
            {
                app.UseSwagger();
                app.UseSwaggerUI();
            }

            app.UseHttpsRedirection();

            app.UseAuthorization();

            app.UseRouting();

            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
            });
        }
    }
}
