using Microsoft.Extensions.Logging;
using Togo.Core.AppServices.TaskItems.Dtos;
using Togo.Core.Entities;
using Togo.Core.Exceptions;
using Togo.Core.Interfaces;
using Togo.Core.Interfaces.AppServices;

namespace Togo.Core.AppServices.TaskItems;

public class TaskItemAppService : ITaskItemAppService
{
    private readonly IAppDbContext _dbContext;
    private readonly ICurrentUserService _currentUserService;
    private readonly ILogger<TaskItemAppService> _logger;

    public TaskItemAppService(
        IAppDbContext dbContext, 
        ICurrentUserService currentUserService, 
        ILogger<TaskItemAppService> logger)
    {
        _dbContext = dbContext;
        _currentUserService = currentUserService;
        _logger = logger;
    }

    public async Task<TaskItemDto> CreateAsync(CreateTaskItemDto input)
    {
        var todayTasksCount = _dbContext.TaskItems.Count(task => 
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
        _dbContext.TaskItems.Add(taskItem);
        await _dbContext.SaveChangesAsync();

        return new TaskItemDto(taskItem);
    }
}
