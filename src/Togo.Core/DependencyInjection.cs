using Microsoft.Extensions.DependencyInjection;
using Togo.Core.AppServices.TaskItems;
using Togo.Core.Interfaces.AppServices;

namespace Togo.Core;

public static class DependencyInjection
{
    public static IServiceCollection AddCoreServices(this IServiceCollection services)
    {
        services.AddAppServices();
        return services;
    }   
    
    private static IServiceCollection AddAppServices(this IServiceCollection services)
    {
        services.AddScoped<ITaskItemAppService, TaskItemAppService>();
        return services;
    }     
}
