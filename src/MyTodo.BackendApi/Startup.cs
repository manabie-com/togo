using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.HttpsPolicy;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Microsoft.IdentityModel.Logging;
using Microsoft.OpenApi.Models;
using MyTodo.Data.Entities;
using MyTodo.Data.EntityFramework;
using MyTodo.Data.EntityFramework.Repositories;
using MyTodo.Data.Interfaces.Repositories;
using MyTodo.Infrastructure.Interfaces;
using MyTodo.Services.Config.AutoMapper;
using MyTodo.Services.Impl;
using MyTodo.Services.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MyTodo.BackendApi
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            //Configire Cors
            services.ConfigureCors();

            //Dependency injection
            services.ConfigureDependencyInjection(Configuration);

            //Add Authentication
            services.Authenticate(Configuration);

            // Auto Mapper
            services.AddSingleton(AutoMapperConfig.RegisterMappings().CreateMapper());

            services.AddControllers().AddJsonOptions(options =>
            {
                options.JsonSerializerOptions.PropertyNamingPolicy = null;
            });

            services.ConfigureSwagger();

            services.AddRouting(options => options.LowercaseUrls = true);

        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
                IdentityModelEventSource.ShowPII = true;


            }
            app.ConfigureSwagger();

            app.UseHttpsRedirection();

            app.UseRouting();

            //Add Authentication
            app.Authenticate();

            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
            });
        }
    }
}
