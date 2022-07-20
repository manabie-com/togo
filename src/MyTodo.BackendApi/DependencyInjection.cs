using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using MyTodo.Data.Entities;
using MyTodo.Data.EntityFramework;
using MyTodo.Data.EntityFramework.Repositories;
using MyTodo.Data.Interfaces.Repositories;
using MyTodo.Infrastructure.Interfaces;
using MyTodo.Services.Impl;
using MyTodo.Services.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MyTodo.BackendApi
{
    public static class DependencyInjection
    {
        public static IServiceCollection ConfigureDependencyInjection(this IServiceCollection services, IConfiguration configuration)
        {
            if (services == null)
            {
                throw new ArgumentNullException(nameof(services));
            }
            services.AddDbContext<MyTodoDbContext>(options =>
                 options.UseSqlServer(configuration.GetConnectionString("DefaultConnection"),
                 o => o.MigrationsAssembly("MyTodo.Data.EntityFramework")));

            // Add application services.
            services.AddScoped<UserManager<AppUser>, UserManager<AppUser>>();
            services.AddScoped<RoleManager<AppRole>, RoleManager<AppRole>>();
            services.AddTransient(typeof(IUnitOfWork), typeof(EFUnitOfWork));
            services.AddTransient(typeof(IRepository<,>), typeof(EFRepository<,>));
            services.AddTransient<ITodoItemRepository, TodoItemRepository>();
            services.AddTransient<IAssignmentRepository, AssignmentRepository>();
            services.AddTransient<ITodoItemService, TodoItemService>();
            services.AddTransient<IAssignmentService, AssignmentService>();
            services.AddTransient<DatabaseInitializer>();
            return services;
        }

    }
}
