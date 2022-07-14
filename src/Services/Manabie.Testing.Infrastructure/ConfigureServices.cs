using Manabie.Testing.Application.Interfaces;
using Manabie.Testing.Infrastructure.Persistance;
using Manabie.Testing.Infrastructure.Persistence;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;

namespace Manabie.Testing.Infrastructure
{
    public static class ConfigureServices
    {
        public static IServiceCollection AddInfrastructure(this IServiceCollection services)
        {
            services.AddDbContext<ManabieDbContext>(options =>
              options.UseInMemoryDatabase("CleanArchitectureDb"));

            services.AddScoped<IManabieDbContext>(provider => provider.GetRequiredService<ManabieDbContext>());
            services.AddScoped<DataSeed>();
            return services;
        }
    }
}
