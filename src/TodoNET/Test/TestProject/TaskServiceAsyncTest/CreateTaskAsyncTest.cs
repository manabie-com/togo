using Application.DTOs.Task;
using Application.Exceptions;
using Infrastructure.Persistence.Services;
using Microsoft.AspNetCore.Http;
using NUnit.Framework;
using System;
using System.Collections.Generic;
using System.Security.Claims;
using TodoNet.Api.Services;

namespace TestProject.TaskServiceAsyncTest
{
    [TestFixture]
    class CreateTaskAsyncTest : TaskServiceAsyncBase
    {
        private TaskServiceAsync taskService;
        private CreateTaskRequest createTaskRequest;
        protected override void RunBeforeTest()
        {
            taskService = new TaskServiceAsync(genericTaskRepositoryAsync, mapper, genericUserRepositoryAsync, dbContext, authenticatedUserService);
            createTaskRequest = new CreateTaskRequest()
            {
                Content = "Content"
            };
        }
        [Test]
        public void Should_Throw_Expected_Exception_When_Request_Is_Null()
        {
            RunTestMethodAndRollback(() =>
            {
                Assert.ThrowsAsync<ArgumentNullException>(async () => await taskService.CreateTaskAsync(null), "Service should throw ArgumentNullException when request is null");
            });
        }
        [Test]
        public void Should_Throw_Expected_Exception_When_Daily_Limit_Is_Reached()
        {
            RunTestMethodAndRollback(() =>
            {
                void Run()
                {
                    try
                    {
                        while (true)
                        {
                            var request = createTaskRequest;
                            taskService.CreateTaskAsync(request).Wait();
                        }
                    }
                    catch (Exception e)
                    {
                        throw e.InnerException;
                    }
                }
                Assert.Throws(Is.TypeOf<ApiException>(), Run);
            });
        }
        [Test]
        public void Should_Throw_Expected_Exception_When_Content_Is_Null_Or_Empty()
        {
            var request = new CreateTaskRequest()
            {
                Content = string.Empty
            };
            RunTestMethodAndRollback(() =>
            {
                Assert.ThrowsAsync<ApiException>(async () => await taskService.CreateTaskAsync(request), "Service should throw ApiException when Content is null or empty");
            });
        }
        [Test]
        public void Should_Throw_Expected_Exception_User_Not_Found()
        {
            RunTestMethodAndRollback(() =>
            {
                var claims = new List<Claim>()
                {
                    new Claim("uid","secondUser")
                };
                var identity = new ClaimsIdentity();
                identity.AddClaims(claims);

                var contextUser = new ClaimsPrincipal(identity);
                HttpContextAccessor httpContextAccessor = new HttpContextAccessor();
                httpContextAccessor.HttpContext = new DefaultHttpContext();
                httpContextAccessor.HttpContext.User = contextUser;
                var errorAuthen = new AuthenticatedUserService(httpContextAccessor);
                var errSerivce = new TaskServiceAsync(genericTaskRepositoryAsync, mapper, genericUserRepositoryAsync, dbContext, errorAuthen);
                Assert.ThrowsAsync<KeyNotFoundException>(async () => await errSerivce.CreateTaskAsync(createTaskRequest), "Service should throw KeyNotFoundException when user not found");
            });
        }
        [Test]
        public void Should_Run_Successfully()
        {
            RunTestMethodAndRollback(() =>
            {
                var request = createTaskRequest;
                var task = taskService.CreateTaskAsync(request);
                var response = task.Result;
                Assert.NotNull(response);
                Assert.NotNull(response.Data);
                Assert.AreEqual(response.Data.Content, request.Content);
                Assert.NotNull(response.Data.Id);
                Assert.NotNull(response.Data.UserId);
                Assert.AreEqual(response.Succeeded, true, "Should return success status");
            });
        }
    }
}
