using Application.DTOs.User;
using Application.Exceptions;
using Infrastructure.Persistence.Services;
using NUnit.Framework;
using System;

namespace TestProject.UserServiceAsyncTest
{
    [TestFixture]
    class AuthenticateUserAsyncTest : UserServiceAsyncBase
    {
        private UserServiceAsync userService;
        private AuthenticationRequest authenRequest;
        protected override void RunBeforeTest()
        {
            userService = new UserServiceAsync(dbContext, jwtOption, appSettings, genericUserRepositoryAsync, mapper);
            authenRequest = new AuthenticationRequest()
            {
               Email = "firstUser@gmail.com",
               Password  = "firstUser"
            };
        }
        [Test]
        public void Should_Throw_Expected_Exception_When_Request_Is_Null()
        {
            RunTestMethodAndRollback(() =>
            {
                Assert.ThrowsAsync<ArgumentNullException>(async () => await userService.AuthenticateAsync(null), "Service should throw ArgumentNullException when request is null");
            });
        }
        [Test]
        public void Should_Run_Successfully()
        {
            RunTestMethodAndRollback(() =>
            {
                var request = authenRequest;
                var task =  userService.AuthenticateAsync(request);
                var response = task.Result;

                Assert.NotNull(response);
                Assert.NotNull(response.Data);
                Assert.AreEqual(response.Succeeded, true, "Should return success status");
                Assert.NotNull(response.Data.JWToken);
                Assert.AreEqual(response.Data.Email, request.Email);
            });
        }
        [Test]
        public void Should_Throw_Expected_Exception_When_Email_Not_Exist()
        {
            var request = new CreateUserRequest()
            {
                Email = "Test@gmail.com",
                Password = "Test"
            };
            RunTestMethodAndRollback(() =>
            {
                Assert.ThrowsAsync<ApiException>(async () => await userService.AuthenticateAsync(request), "Service should throw ApiException when request is null");
            });
        }
    }
}
