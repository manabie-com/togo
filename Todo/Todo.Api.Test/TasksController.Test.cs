using Microsoft.AspNetCore.Mvc;
using Moq;
using System;
using System.Net;
using System.Threading.Tasks;
using Todo.Api.Controllers;
using Todo.Application.Dtos;
using Todo.Application.Services;
using Xunit;

namespace Todo.Api.Test
{
    public class TasksControllerTest
    {
        [Fact]
        public void TaskService_Create_Task_Success()
        {
            //Arrange
            var input = new CreateEditTaskDto()
            {
                Title = "This is test",
                Description = "This is description test",
                UserId = Guid.NewGuid().ToString(),
                Type = 0,
                Priority = 0,
            };
            var _mockTaskService = new Mock<IUserTaskService>();
            _mockTaskService.Setup(x => x.CreateTaskAsync(input)).Returns(Task.FromResult("success"));

            //Act

            TasksController controller = new TasksController(_mockTaskService.Object);
            var result = controller.CreateTask(input).Result as ObjectResult;

            //Assert

            Assert.Equal(HttpStatusCode.OK, (HttpStatusCode)result.StatusCode);
        }
    }
}