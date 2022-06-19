using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Npgsql;
using Todo.Application.Interfaces;

namespace Todo.Infrastructure
{
    public static class ServiceConfiguration
    {
        public static IServiceCollection AddInfrastructure(this IServiceCollection services, IConfiguration configuration)
        {
            string connectionString = configuration.GetConnectionString("DefaultConnection");

            if (string.IsNullOrEmpty(connectionString))
            {
                throw new ArgumentNullException(nameof(connectionString), "Connection string should not empty.");
            }

            services.AddDbContext<ApplicationDbContext>(options => options.UseNpgsql(connectionString, configs => configs.MigrationsHistoryTable("__EFMigrationsHistory", "public")));

            services.AddScoped<IApplicationDbContext, ApplicationDbContext>();

            services.AddTransient<ApplicationDbSeeder>();

            AppContext.SetSwitch("Npgsql.EnableLegacyTimestampBehavior", true);

            return services;
        }
    }
}
