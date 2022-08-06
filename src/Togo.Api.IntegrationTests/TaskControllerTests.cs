using System;
using System.Linq;
using System.Net;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc.Testing;
using Newtonsoft.Json;
using Togo.Core.AppServices.TaskItems.Dtos;
using Togo.Infrastructure.Identities.Dtos;
using Xunit;

namespace Togo.Api.IntegrationTests;

public class TaskControllerTests : BaseIntegrationTestClass
{
    private const string CreateTaskPath = "api/tasks";
    
    public TaskControllerTests(WebApplicationFactory<Program> webApplicationFactory) : base(webApplicationFactory)
    {
        LoginAs(AdminUserName, AdminPassword).Wait();
    }

    [Theory]
    [InlineData(null)]
    [InlineData("")]
    public async Task CreateAsync_WhenTitleIsNullOrEmpty_ShouldReturnBadRequest(string title)
    {
        var createTaskInput = new CreateTaskItemDto { Title = title };
        var response = await PostAsync(CreateTaskPath, createTaskInput);
        
        Assert.Equal(HttpStatusCode.BadRequest, response.StatusCode);
    }

    [Fact]
    public async Task CreateAsync_WhenTitleIsNotNullOrEmpty_ShouldSuccess()
    {
        const string testTitle = "Discuss with Kelven about new project";
        var createTaskInput = new CreateTaskItemDto { Title = testTitle };
        var response = await PostAsync(CreateTaskPath, createTaskInput);
        
        Assert.Equal(HttpStatusCode.OK, response.StatusCode);

        var responseString = await response.Content.ReadAsStringAsync();
        var responseObject = JsonConvert.DeserializeObject<TaskItemDto>(responseString);
        
        Assert.NotEqual(0, responseObject.Id);
        Assert.Equal(testTitle, responseObject.Title);
    }

    [Fact]
    public async Task CreateAsync_WhenTaskLimitExceeded_ShouldReturnBadRequest()
    {
        var testUserName = Guid.NewGuid().ToString();
        var testPassword = "Abcd@1234";
        var testMaxTasksPerDay = 5;

        var createUserInput = new CreateUserDto
        {
            UserName = testUserName,
            Password = testPassword,
            MaxTasksPerDay = testMaxTasksPerDay
        };

        await PostAsync(CreateUserPath, createUserInput);
        await LoginAs(testUserName, testPassword);

        var createTaskInputs = Enumerable.Range(0, 5)
            .Select(number => new CreateTaskItemDto { Title = $"Task {number}" });

        var creationTasks = createTaskInputs
            .Select(createTaskInput => PostAsync(CreateTaskPath, createTaskInput))
            .ToList();

        await Task.WhenAll(creationTasks);

        var createTaskInput = new CreateTaskItemDto { Title = "Last task" };
        var response = await PostAsync(CreateTaskPath, createTaskInput);
        
        Assert.Equal(HttpStatusCode.BadRequest, response.StatusCode);
    }
}
