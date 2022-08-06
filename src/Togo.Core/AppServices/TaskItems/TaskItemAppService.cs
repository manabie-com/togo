using Microsoft.Extensions.Logging;
using Togo.Core.AppServices.TaskItems.Dtos;
using Togo.Core.Entities;
using Togo.Core.Exceptions;
using Togo.Core.Interfaces;
using Togo.Core.Interfaces.AppServices;

namespace Togo.Core.AppServices.TaskItems;

public class TaskItemAppService : ITaskItemAppService
{
    private readonly ICurrentUserService _currentUserService;
    private readonly ILogger<TaskItemAppService> _logger;
    private readonly IUnitOfWork _unitOfWork;

    public TaskItemAppService(
        ICurrentUserService currentUserService, 
        ILogger<TaskItemAppService> logger, 
        IUnitOfWork unitOfWork)
    {
        _currentUserService = currentUserService;
        _logger = logger;
        _unitOfWork = unitOfWork;
    }

    public async Task<TaskItemDto> CreateAsync(CreateTaskItemDto input)
    {
        var todayTasksCount = await _unitOfWork.TaskItemRepository
            .CountAsync(task =>
                task.CreatedAt.Date.Equals(DateTime.UtcNow.Date)
                && task.CreatedBy == _currentUserService.UserId);

        if (todayTasksCount >= _currentUserService.MaxTasksPerDay)
        {
            _logger.LogWarning(
                "User {UserId} reached task limit of {TaskLimit} on {Today} and cannot create more", 
                _currentUserService.UserId,
                _currentUserService.MaxTasksPerDay,
                DateTimeOffset.UtcNow.Date);

            throw new TaskLimitExceededException(_currentUserService.MaxTasksPerDay);
        }
        
        var taskItem = TaskItem.CreateNew(input.Title);
        await _unitOfWork.TaskItemRepository.AddAsync(taskItem);

        await _unitOfWork.CommitAsync();

        return new TaskItemDto(taskItem);
    }
}
