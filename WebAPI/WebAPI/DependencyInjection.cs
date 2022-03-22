using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Models;
using Repositories.Infrastructure;
using Services.Implementations;
using Services.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace WebAPI
{
    public class DependencyInjection
    {
        public static void Start(IConfiguration configuration, IServiceCollection services)
        {
            services.AddSingleton<IHttpContextAccessor, HttpContextAccessor>();

            services.AddTransient<IUnitOfWork, UnitOfWork>();

            services.AddTransient<IRepository<Users>, Repository<Users>>();
            services.AddTransient<IUserService, UserService>();

            services.AddTransient<IRepository<Tasks>, Repository<Tasks>>();
            services.AddTransient<ITaskService, TaskService>();
        }
    }
}
