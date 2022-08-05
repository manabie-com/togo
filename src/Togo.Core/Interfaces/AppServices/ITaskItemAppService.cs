using Togo.Core.AppServices.TaskItems.Dtos;

namespace Togo.Core.Interfaces.AppServices;

public interface ITaskItemAppService
{
    Task<TaskItemDto> CreateAsync(CreateTaskItemDto input);
}
