using MockQueryable.Moq;
using Moq;
using System;
using WebApi.Models;
using WebApi.Requests;
using WebApi.Services;
using Xunit;

namespace WebApi.Test
{
    public class TaskServiceTest
    {
        private readonly TaskService _taskService;
        private readonly Mock<DemoDbContext> _mockDbContext;

        public TaskServiceTest()
        {
            _mockDbContext = new Mock<DemoDbContext>();
            _taskService = new TaskService(_mockDbContext.Object);
        }

        [Theory]
        [InlineData("Task 001", "ee08f09c-319c-484c-936f-0c020e343bf5")]
        [InlineData("Task 002", "e75b9289-0ed6-45a2-96c6-b12848f958c6")]
        public async System.Threading.Tasks.Task CreateTask_Success(string taskContent, string userId)
        {
            // Setup
            var createTaskRequest = new CreateTaskRequest
            {
                Content = taskContent
            };

            var result = await _taskService.CreateTask(createTaskRequest, Guid.Parse(userId));

            // Assert
            Assert.NotNull(result);
            Assert.Equal(result.Content, taskContent);
        }

        [Theory]
        [InlineData("12b8d928-39ca-4c5b-a01f-9b24f917deb2", "ee08f09c-319c-484c-936f-0c020e343bf5|2429507e-d057-44cb-9a6a-cce9447199a7|e75b9289-0ed6-45a2-96c6-b12848f958c6|bc347f5c-5c0b-4157-9d37-8645a3b11302|f2e120e9-1fd0-43a1-92f1-651a389a972a", null, 1)]
        [InlineData("ee08f09c-319c-484c-936f-0c020e343bf5", "ee08f09c-319c-484c-936f-0c020e343bf5|2429507e-d057-44cb-9a6a-cce9447199a7|e75b9289-0ed6-45a2-96c6-b12848f958c6|bc347f5c-5c0b-4157-9d37-8645a3b11302|f2e120e9-1fd0-43a1-92f1-651a389a972a", true, 3)]
        [InlineData("ee08f09c-319c-484c-936f-0c020e343bf5", "ee08f09c-319c-484c-936f-0c020e343bf5|2429507e-d057-44cb-9a6a-cce9447199a7|e75b9289-0ed6-45a2-96c6-b12848f958c6|bc347f5c-5c0b-4157-9d37-8645a3b11302|f2e120e9-1fd0-43a1-92f1-651a389a972a", true, 4)]
        [InlineData("ee08f09c-319c-484c-936f-0c020e343bf5", "ee08f09c-319c-484c-936f-0c020e343bf5|2429507e-d057-44cb-9a6a-cce9447199a7|e75b9289-0ed6-45a2-96c6-b12848f958c6|bc347f5c-5c0b-4157-9d37-8645a3b11302|f2e120e9-1fd0-43a1-92f1-651a389a972a", false, 5)]
        public async System.Threading.Tasks.Task CheckLimitDailyTask(string userId, string userIds, bool? expectedResult, int totalTaskRecords)
        {
            // Setup
            var users = TestServiceUtilities.SetupDataForUserEntity(userIds, "123456");
            var tasks = TestServiceUtilities.SetupDataForTaskEntity(Guid.Parse(userId), totalTaskRecords);
            var usersMock = users.BuildMockDbSet();
            var tasksMock = tasks.BuildMockDbSet();
            _mockDbContext.Setup(x => x.Users).Returns(usersMock.Object);
            _mockDbContext.Setup(x => x.Tasks).Returns(tasksMock.Object);

            var result = await _taskService.CheckLimitDailyTask(Guid.Parse(userId));

            // Assert
            Assert.Equal(expectedResult, result);
        }
    }
}
