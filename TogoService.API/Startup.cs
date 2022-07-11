using AutoMapper;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Microsoft.OpenApi.Models;
using TogoService.API.Filter;
using TogoService.API.Infrastructure.Database;
using TogoService.API.Infrastructure.Repository;
using TogoService.API.Model.Interface;
using System;
using System.IO;
using System.Reflection;
using Microsoft.EntityFrameworkCore.Diagnostics;

namespace TogoService.API
{
    public class Startup
    {
        string swaggerDocVersion = "v1";
        string swaggerDocTitle = "Togo Service Api";
        ILoggerFactory _loggerFactory = new LoggerFactory();
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            // var writerConnectionStr = Environment.GetEnvironmentVariable(EnvironmentVariableNames.EnvNameConnStringWriter);
            // var readerConnectionStr = Environment.GetEnvironmentVariable(EnvironmentVariableNames.EnvNameConnStringReader);

            services.AddDbContext<TogoDbContext>(options => options.UseLazyLoadingProxies()
                .ConfigureWarnings(b => b.Log((RelationalEventId.CommandCreated, LogLevel.Information),
                    (RelationalEventId.CommandExecuted, LogLevel.Information),
                    (RelationalEventId.TransactionCommitted, LogLevel.Information)))
                .UseSqlite("Data Source=TogoService.db;")
            );

            // Repository - Unit of Work
            services.AddScoped<IUnitOfWork, UnitOfWork>();            
            services.AddScoped<ITodoTaskRepository, TodoTaskRepository>();

            services.AddControllers();

            services.AddAutoMapper(typeof(Startup));

            services.AddLogging(logging =>
            {
                logging.AddConsole();
                logging.SetMinimumLevel(LogLevel.Information);
            });

            services.AddCors(c => c.AddPolicy("AllowOrigin", o => o.AllowAnyOrigin()));
            services.AddControllers()
                .AddJsonOptions(options =>
                options.JsonSerializerOptions.Converters.Add(new System.Text.Json.Serialization.JsonStringEnumConverter()));
            services.AddScoped<ValidationActionFilter>();

            services.AddSwaggerGen(options =>
            {
                options.SwaggerDoc(swaggerDocVersion, new OpenApiInfo { Title = swaggerDocTitle, Version = swaggerDocVersion });
                var xmlFile = $"{Assembly.GetExecutingAssembly().GetName().Name}.xml";
                var xmlPath = Path.Combine(AppContext.BaseDirectory, xmlFile);
                options.IncludeXmlComments(xmlPath);
                options.CustomSchemaIds(c => c.FullName);

                options.AddSecurityDefinition("oauth2", new OpenApiSecurityScheme
                {
                    Description = "Standard Authorization header using the Bearer scheme. Example: \"bearer {token}\"",
                    Name = "Authorization",
                    In = ParameterLocation.Header,
                    Type = SecuritySchemeType.ApiKey,
                    Scheme = "Bearer"
                });
            });
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            // Apply EF migration to database
            using (var scope = app.ApplicationServices.CreateScope())
            {
                using var context = scope.ServiceProvider.GetService<TogoDbContext>();

                for (int i = 1; i <= 5; i++)
                {
                    try
                    {
                        Console.WriteLine($"Try execute migrate command time: {i}");
                        context.Database.Migrate();
                        break;
                    }
                    catch (Exception ex)
                    {
                        Console.WriteLine(ex.Message);
                    }
                }
            }

            app.UseExceptionHandler("/api/error");

            app.UseHttpsRedirection();

            app.UseSwagger();
            app.UseSwaggerUI(c =>
            {
                c.SwaggerEndpoint("/swagger/v1/swagger.json", swaggerDocTitle);
            });

            app.UseRouting();

            app.UseAuthentication();
            app.UseAuthorization();
            app.UseCors(o => o.AllowAnyOrigin());
            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
            });
        }
    }
}