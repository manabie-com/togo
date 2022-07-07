using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using TogoService.API.Dto;
using TogoService.API.Infrastructure.Helper.MessageUtil;
using TogoService.API.Model;
using TogoService.IntegrationTest.Helper;
using Xunit;

namespace TogoService.IntegrationTest.Tests
{
    public partial class UserControllerTest
    {
        [Fact]
        public async Task Successful_201AddedSuccessfully()
        {
            // Arrange            
            NewTaskRequest request = FakeData.GenerateNewTaskRequest(2);
            var expectedResponse = new CommonResponse<string>(StatusCodes.Status201Created,
                CommonMessages.Ok,
                CommonMessages.GetSuccessfulAddedItemsMsg(typeof(TodoTask).Name, request.Tasks.Length));

            // Act
            var result = (ObjectResult)await _userController.AddTasksForUser(_userWith10MaxDailyTasks.Id, request);

            // Assert
            Assert.NotNull(result);
            Assert.Equal(StatusCodes.Status201Created, result.StatusCode);
            CommonResponse<string> actualResponse = (CommonResponse<string>)result.Value;
            Assert.Equal(expectedResponse.StatusCode, actualResponse.StatusCode);
            Assert.Equal(expectedResponse.Message, actualResponse.Message);
            Assert.Equal(expectedResponse.Result, actualResponse.Result);
        }

        [Fact]
        public async Task UserIdNotFound_404NotFound()
        {
            // Arrange
            Guid notFoundUserId = Guid.NewGuid();
            var expectedResponse = new CommonResponse<string>(StatusCodes.Status404NotFound,
                CommonMessages.GetCannotFindMsg(typeof(API.Model.User).Name, notFoundUserId),
                null);

            // Act
            var result = (ObjectResult)await _userController.AddTasksForUser(notFoundUserId, FakeData.GenerateNewTaskRequest(1));

            // Assert
            Assert.NotNull(result);
            Assert.Equal(StatusCodes.Status404NotFound, result.StatusCode);
            CommonResponse<string> actualResponse = (CommonResponse<string>)result.Value;
            Assert.Equal(expectedResponse.StatusCode, actualResponse.StatusCode);
            Assert.Equal(expectedResponse.Message, actualResponse.Message);
        }

        [Fact]
        public async Task AddMoreThanMaxTasksPerDay_422ReachoutMaxTaskPerDay()
        {
            // Arrange
            NewTaskRequest request = FakeData.GenerateNewTaskRequest(2);
            var expectedResponse = new CommonResponse<string>(StatusCodes.Status422UnprocessableEntity,
                UserControllerErrMsg.ReachoutMaxTaskPerDay,
                null);

            // Act
            var result = (ObjectResult)await _userController.AddTasksForUser(_userWith0MaxDailyTasks.Id, request);

            // Assert
            Assert.NotNull(result);
            Assert.Equal(StatusCodes.Status422UnprocessableEntity, result.StatusCode);
            CommonResponse<string> actualResponse = (CommonResponse<string>)result.Value;
            Assert.Equal(expectedResponse.StatusCode, actualResponse.StatusCode);
            Assert.Equal(expectedResponse.Message, actualResponse.Message);
        }
    }
}