using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Todo.Application.Mapping;
using Todo.Application.Services;

namespace Todo.Application
{
    public static class ServiceConfiguration
    {
        public static IServiceCollection AddApplication(this IServiceCollection services, IConfiguration configuration)
        {

            var allowedOrigins = configuration["AllowOrigins"]?.Split(new[] { ';' }, StringSplitOptions.RemoveEmptyEntries);


            //services.AddCors(options => options.AddPolicy("TodoPolicy", policy =>
            //{
            //    policy.AllowAnyMethod().AllowAnyHeader().AllowAnyOrigin();
            //}));
            services.AddMapping();
            return services.AddAppServices(configuration);
        }

        public static IServiceCollection AddAppServices(this IServiceCollection services, IConfiguration configuration)
        {
            services.AddTransient<IUserTaskService, UserTaskService>();

            return services;
        }
    }
}
