using Microsoft.Extensions.DependencyInjection;
using Models;
using Services.Interfaces;
using Xunit;

namespace IntergationTest
{
    [Collection("Sequential")]
    [TestCaseOrderer("IntergationTest.Config.PriorityOrderer", "IntergationTest")]
    public class TaskServiceTest : BaseService_Test
    {
        private ITaskService _taskService;
        public TaskServiceTest() : base()
        {
            _taskService = ServiceProvider.GetRequiredService<ITaskService>();
        }

        [Fact]
        public void Should_Create_Success()
        {
            var task = new Tasks()
            {
                Content = "Task Content"
            };

            var result = _taskService.Create(task, "1");

            Assert.Equal(1, result);
        }

        [Fact]
        public void Should_Create_Fail()
        {
            var task = new Tasks()
            {
                Content = "Task Content",
            };

            var result = _taskService.Create(task, "2");

            Assert.Equal(0, result);
        }

        [Fact]
        public void Should_Create_Exception()
        {
            var task = new Tasks()
            {
                Content = "Task Content",
            };

            var result = _taskService.Create(task, "5");

            Assert.Equal(-1, result);
        }

        [Fact]
        public void Should_GetTaskByUserId_HaveData()
        {
            var result = _taskService.GetTasksByUserId("2");

            Assert.True(result.Count > 0);
        }

        [Fact]
        public void Should_GetTaskByUserId_HaveNoData()
        {
            var result = _taskService.GetTasksByUserId("3");

            Assert.True(result.Count == 0);
        }
    }
}
