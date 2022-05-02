using Manabie.Togo.Api.Configurations;
using Manabie.Togo.Core.Bus;
using Manabie.Togo.Domain.Commands.UserTask.Create;
using Manabie.Togo.JsonRepository.UserTask;
using Manabie.Togo.RedisRepository.Implememt;
using Manabie.Togo.RedisRepository.Interface;
using Manabie.Togo.Service.Implememt.UserTask;
using Manabie.Togo.Service.Interface.UserTask;
using MediatR;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.HttpsPolicy;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Microsoft.OpenApi.Models;
using StackExchange.Redis;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;
using System.Text.Json.Serialization;
using System.Threading.Tasks;

namespace Manabie.Togo.Api
{
	public class Startup
	{
		public Startup(IConfiguration configuration)
		{
			_config = configuration;
		}

		private readonly IConfiguration _config;

		// This method gets called by the runtime. Use this method to add services to the container.
		public void ConfigureServices(IServiceCollection services)
		{

			services.AddControllers();
			services.AddSwaggerGen(c =>
			{
				c.SwaggerDoc("v1", new OpenApiInfo { Title = "Manabie.Togo.Api", Version = "v1" });
			});

			// Auto maper
			services.AddCommonServices();

			// Json reading file
			services.AddSingleton<IUserTaskJsonRepository, UserTaskJsonRepository>();

			// Mediator
			services.AddMediatR(typeof(CreatedUserTaskCommand).GetTypeInfo().Assembly);
			services.AddScoped<IMediatorHandler, InMemoryBus>();

			// Service
			services.AddScoped<IUserTaskService, UserTaskService>();

			// Redis Repository
			services.AddSingleton<IUserTaskRepositoryRedis, UserTaskRepositoryRedis>();

			// Init Redis
			string redisConnection = _config.GetSection("MyConfigs")["RedisConnection"];
			services.AddSingleton<IConnectionMultiplexer>(ConnectionMultiplexer.Connect(redisConnection));
			services.AddSingleton<IDatabase>(sp =>
			{
				var con = sp.GetService<IConnectionMultiplexer>();
				return con.GetDatabase();
			});
		}

		// This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
		public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
		{
			if (env.IsDevelopment())
			{
				app.UseDeveloperExceptionPage();
				app.UseSwagger();
				app.UseSwaggerUI(c => c.SwaggerEndpoint("/swagger/v1/swagger.json", "Manabie.Togo.Api v1"));
			}

			app.UseHttpsRedirection();

			app.UseRouting();

			app.UseAuthorization();

			app.UseEndpoints(endpoints =>
			{
				endpoints.MapControllers();
			});
		}
	}
}
