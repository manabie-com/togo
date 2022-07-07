using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Moq;
using TogoService.API.Dto;
using TogoService.API.Infrastructure.Helper.MessageUtil;
using TogoService.API.Model;
using TogoService.UnitTest.Helper;
using Xunit;

namespace TogoService.UnitTest.Tests
{
    public partial class UserControllerTest
    {
        private Guid _emptyGuid = Guid.Empty;

        [Fact]
        public async Task NullRequest_ThrowException()
        {
            // Act
            var exception = await Record.ExceptionAsync(async () => await _userController.AddTasksForUser(_emptyGuid, null));

            // Assert
            Assert.NotNull(exception);
        }

        [Fact]
        public async Task MissingTodoDate_404NotFound()
        {
            // Arrange
            Guid userId = Guid.NewGuid();
            var expectedResponse = new CommonResponse<string>(StatusCodes.Status400BadRequest,
                UserControllerErrMsg.MissingTodoDay,
                null);
            NewTaskRequest request = FakeData.GenerateNewTaskRequest();
            request.Date = new DateTime();

            // Act
            var result = (ObjectResult)await _userController.AddTasksForUser(userId, request);

            // Assert
            Assert.NotNull(result);
            Assert.Equal(StatusCodes.Status400BadRequest, result.StatusCode);
            CommonResponse<string> actualResponse = (CommonResponse<string>)result.Value;
            Assert.Equal(expectedResponse.StatusCode, actualResponse.StatusCode);
            Assert.Equal(expectedResponse.Message, actualResponse.Message);
        }

        [Fact]
        public async Task EmptyUserId_404NotFound()
        {
            // Arrange
            var expectedResponse = new CommonResponse<string>(StatusCodes.Status404NotFound,
                CommonMessages.GetCannotFindMsg(typeof(API.Model.User).Name, _emptyGuid),
                null);

            // Act
            var result = (ObjectResult)await _userController.AddTasksForUser(_emptyGuid, FakeData.GenerateNewTaskRequest());

            // Assert
            Assert.NotNull(result);
            Assert.Equal(StatusCodes.Status404NotFound, result.StatusCode);
            CommonResponse<string> actualResponse = (CommonResponse<string>)result.Value;
            Assert.Equal(expectedResponse.StatusCode, actualResponse.StatusCode);
            Assert.Equal(expectedResponse.Message, actualResponse.Message);
        }

        [Fact]
        public async Task UserIdNotFound_404NotFound()
        {
            // Arrange
            Guid userId = Guid.NewGuid();
            _mockGenericUserRepository.Setup(q => q.GetById(userId)).ReturnsAsync((API.Model.User)null);
            var expectedResponse = new CommonResponse<string>(StatusCodes.Status404NotFound,
                CommonMessages.GetCannotFindMsg(typeof(API.Model.User).Name, userId),
                null);

            // Act
            var result = (ObjectResult)await _userController.AddTasksForUser(userId, FakeData.GenerateNewTaskRequest());

            // Assert
            Assert.NotNull(result);
            Assert.Equal(StatusCodes.Status404NotFound, result.StatusCode);
            CommonResponse<string> actualResponse = (CommonResponse<string>)result.Value;
            Assert.Equal(expectedResponse.StatusCode, actualResponse.StatusCode);
            Assert.Equal(expectedResponse.Message, actualResponse.Message);
        }

        [Fact]
        public async Task UnexpectedException_500InternalServerError()
        {
            // Arrange
            Guid userId = Guid.NewGuid();
            string expectedExceptionMsg = "Exception";
            _mockGenericUserRepository.Setup(q => q.GetById(userId)).ThrowsAsync(new Exception(expectedExceptionMsg));
            var expectedResponse = new CommonResponse<string>(StatusCodes.Status500InternalServerError,
                expectedExceptionMsg,
                null);

            // Act            
            var result = (ObjectResult)await _userController.AddTasksForUser(userId, FakeData.GenerateNewTaskRequest());

            // Assert
            Assert.NotNull(result);
            Assert.Equal(StatusCodes.Status500InternalServerError, result.StatusCode);
            CommonResponse<string> actualResponse = (CommonResponse<string>)result.Value;
            Assert.Equal(expectedResponse.StatusCode, actualResponse.StatusCode);
            Assert.Equal(expectedResponse.Message, actualResponse.Message);
        }

        [Fact]
        public async Task AddMoreThanMaxTasksPerDay_422ReachoutMaxTaskPerDay()
        {
            // Arrange
            User user = FakeData.GenerateUser(10);
            NewTaskRequest request = FakeData.GenerateNewTaskRequest();
            TodoTask[] addedTasks = FakeData.GenerateTodoTasks((int)user.MaxDailyTasks);
            _mockGenericUserRepository.Setup(q => q.GetById(user.Id)).ReturnsAsync(user);
            _mockTodoTaskRepository.Setup(q => q.GetAddedTasks(user.Id, request.Date)).ReturnsAsync(addedTasks);
            var expectedResponse = new CommonResponse<string>(StatusCodes.Status422UnprocessableEntity,
                UserControllerErrMsg.ReachoutMaxTaskPerDay,
                null);

            // Act
            var result = (ObjectResult)await _userController.AddTasksForUser(user.Id, request);

            // Assert
            Assert.NotNull(result);
            Assert.Equal(StatusCodes.Status422UnprocessableEntity, result.StatusCode);
            CommonResponse<string> actualResponse = (CommonResponse<string>)result.Value;
            Assert.Equal(expectedResponse.StatusCode, actualResponse.StatusCode);
            Assert.Equal(expectedResponse.Message, actualResponse.Message);
        }

        [Fact]
        public async Task Successful_201AddedSuccessfully()
        {
            // Arrange
            User user = FakeData.GenerateUser(10);
            NewTaskRequest request = FakeData.GenerateNewTaskRequest(4);
            TodoTask[] addedTasks = FakeData.GenerateTodoTasks(1);
            _mockGenericUserRepository.Setup(q => q.GetById(user.Id)).ReturnsAsync(user);
            _mockTodoTaskRepository.Setup(q => q.GetAddedTasks(user.Id, request.Date)).ReturnsAsync(addedTasks);
            var expectedResponse = new CommonResponse<string>(StatusCodes.Status201Created,
                CommonMessages.Ok,
                CommonMessages.GetSuccessfulAddedItemsMsg(typeof(TodoTask).Name, request.Tasks.Length));

            // Act
            var result = (ObjectResult)await _userController.AddTasksForUser(user.Id, request);

            // Assert
            Assert.NotNull(result);
            Assert.Equal(StatusCodes.Status201Created, result.StatusCode);
            CommonResponse<string> actualResponse = (CommonResponse<string>)result.Value;
            Assert.Equal(expectedResponse.StatusCode, actualResponse.StatusCode);
            Assert.Equal(expectedResponse.Message, actualResponse.Message);
            Assert.Equal(expectedResponse.Result, actualResponse.Result);
        }
    }
}