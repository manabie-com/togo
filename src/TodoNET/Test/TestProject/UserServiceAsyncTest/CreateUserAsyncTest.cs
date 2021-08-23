using Application.DTOs.User;
using Application.Exceptions;
using Infrastructure.Persistence.Services;
using NUnit.Framework;
using System;
using System.Threading.Tasks;

namespace TestProject.UserServiceAsyncTest
{
    [TestFixture]
    class CreateUserAsyncTest : UserServiceAsyncBase
    {
        private UserServiceAsync userService;
        private CreateUserRequest createUserRequest;
        protected override void RunBeforeTest()
        {
            userService = new UserServiceAsync(dbContext, jwtOption, appSettings, genericUserRepositoryAsync, mapper);
            createUserRequest = new CreateUserRequest()
            {
                Email = "secondEmail@gmail.com",
                MaxTodo = 5,
                Password = "secondEmail"
            };
        }
        [Test]
        public void Should_Throw_Expected_Exception_When_Request_Is_Null()
        {
            RunTestMethodAndRollback(() =>
            {
                Assert.ThrowsAsync<ArgumentNullException>(async () => await userService.CreateUserAsync(null), "Service should throw ArgumentNullException when request is null");
            });
        }
        [Test]
        public void Should_Run_Successfully()
        {
            RunTestMethodAndRollback(() =>
            {
                var request = createUserRequest;
                var task =  Task.Run(async() => await userService.CreateUserAsync(request));
                var response = task.Result;
                Assert.NotNull(response);
                Assert.AreEqual(response.Succeeded, true, "Should return success status");
            });
        }
        [Test]
        public void Should_Return_Error_Message_When_New_Email_Exist()
        {
            var existEmail = new CreateUserRequest()
            {
                Email = "firstUser@gmail.com",
                MaxTodo = 5,
                Password = "secondEmail"
            };
            RunTestMethodAndRollback(() =>
            {
                Assert.ThrowsAsync<ApiException>(async () => await userService.CreateUserAsync(existEmail), "Service should throw ApiException when request is null");
            });
        }
    }
}
