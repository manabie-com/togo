using Microsoft.Extensions.DependencyInjection;
using System.Reflection;

namespace Todo.Application.Mapping
{
    public static class MappingConfig
    {
        public static IServiceCollection AddMapping(this IServiceCollection services)
        {
            return services.AddAutoMapper(Assembly.GetExecutingAssembly());
        }
    }
}
