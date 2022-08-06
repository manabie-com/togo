using System;
using System.Linq.Expressions;
using System.Threading.Tasks;
using Microsoft.Extensions.Logging;
using Moq;
using Togo.Core.AppServices.TaskItems;
using Togo.Core.AppServices.TaskItems.Dtos;
using Togo.Core.Entities;
using Togo.Core.Exceptions;
using Togo.Core.Interfaces;
using Togo.Core.Interfaces.Repositories;
using Xunit;

namespace Togo.Core.Tests.AppServices;

public class TaskItemAppServiceTests
{
    [Theory]
    [InlineData(5, 5)]
    [InlineData(5, 6)]
    public async Task CreateAsync_WhenDailyLimitReached_ShouldThrowTaskLimitExceededException(int maxTasksPerDay, int currentNumberOfTasks)
    {
        var mockTaskItemRepository = new Mock<IRepository<TaskItem>>();

        mockTaskItemRepository
            .Setup(repository => repository.CountAsync(It.IsAny<Expression<Func<TaskItem, bool>>>()))
            .ReturnsAsync(currentNumberOfTasks);
        
        var mockUnitOfWork = new Mock<IUnitOfWork>();
        
        mockUnitOfWork
            .Setup(uow => uow.TaskItemRepository)
            .Returns(mockTaskItemRepository.Object);

        var mockCurrentUserService = new Mock<ICurrentUserService>();
        
        mockCurrentUserService
            .Setup(service => service.MaxTasksPerDay)
            .Returns(maxTasksPerDay);

        var mockLogger = new Mock<ILogger<TaskItemAppService>>();

        var taskItemAppService = new TaskItemAppService(mockCurrentUserService.Object, mockLogger.Object, mockUnitOfWork.Object);
        
        await Assert.ThrowsAsync<TaskLimitExceededException>(async () => 
            await taskItemAppService.CreateAsync(new CreateTaskItemDto { Title = "Test task" }));
    }

    [Theory]
    [InlineData(5, 0)]
    [InlineData(5, 1)]
    [InlineData(5, 4)]
    public async Task CreateAsync_WhenDailyLimitNotReached_ShouldSuccess(int maxTasksPerDay, int currentNumberOfTasks)
    {
        var mockTaskItemRepository = new Mock<IRepository<TaskItem>>();

        mockTaskItemRepository
            .Setup(repository => repository.CountAsync(It.IsAny<Expression<Func<TaskItem, bool>>>()))
            .ReturnsAsync(currentNumberOfTasks);
        
        var mockUnitOfWork = new Mock<IUnitOfWork>();
        
        mockUnitOfWork
            .Setup(uow => uow.TaskItemRepository)
            .Returns(mockTaskItemRepository.Object);

        var mockCurrentUserService = new Mock<ICurrentUserService>();
        
        mockCurrentUserService
            .Setup(service => service.MaxTasksPerDay)
            .Returns(maxTasksPerDay);

        var mockLogger = new Mock<ILogger<TaskItemAppService>>();

        var taskItemAppService = new TaskItemAppService(mockCurrentUserService.Object, mockLogger.Object, mockUnitOfWork.Object);

        const string testTaskTitle = "Test task";
        var task = await taskItemAppService.CreateAsync(new CreateTaskItemDto { Title = testTaskTitle });
        
        Assert.Equal(testTaskTitle, task.Title);
    }
}
