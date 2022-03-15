using System;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using FluentAssertions;
using Microsoft.AspNetCore.Mvc.Testing;
using Newtonsoft.Json;
using Todo.Contract;
using Todo.Domain.Models;
using Todo.Storage.Contract;
using Xunit;

namespace Todo.Test.Integration;

public class TodoAppControllerTests : IClassFixture<WebApplicationFactory<Program>>
{
    private readonly HttpClient _client;

    private readonly WebApplicationFactory<Program> _factory;

    private long _userId;

    private long _taskId;

    public TodoAppControllerTests(WebApplicationFactory<Program> factory)
    {
        _factory = factory ?? throw new ArgumentNullException();
        _client = factory.CreateClient();
    }

    [Theory]
    [InlineData("/api/todos")]
    [InlineData("/api/users")]
    public async Task GetEndpointsWithoutQueries_ReturnBadRequest(string url)
    {
        var response = await _client.GetAsync(url);

        response.StatusCode.Should().Be(System.Net.HttpStatusCode.BadRequest);
    }

    [Fact]
    public async Task Post_CreatesTodo_Returns201()
    {
        // Create user
        var userContent = new StringContent(JsonConvert.SerializeObject(CreateSampleUser()), Encoding.UTF8, "application/json");
        var responseCreateUser = await _client.PostAsync($"/api/users", userContent);
        if (responseCreateUser != null)
        {
            _userId = (JsonConvert.DeserializeObject<User>(await responseCreateUser.Content.ReadAsStringAsync())!).Id;
        }

        // Check created user
        var responseGetUser = await _client.GetAsync($"/api/users?id={_userId}");

        // Create todo
        var todoContent = new StringContent(JsonConvert.SerializeObject(CreateSampleTodoRequest(_userId)), Encoding.UTF8, "application/json");
        var responseTodo = await _client.PostAsync($"/api/todos", todoContent);
        if (responseTodo != null)
        {
            _taskId = (JsonConvert.DeserializeObject<Todo.Domain.Models.Todo>(await responseTodo.Content.ReadAsStringAsync())!).Id;
        }

        responseCreateUser!.StatusCode.Should().Be(System.Net.HttpStatusCode.Created);
        responseGetUser.StatusCode.Should().Be(System.Net.HttpStatusCode.OK);
        responseTodo!.StatusCode.Should().Be(System.Net.HttpStatusCode.Created);

        // Dispose created user and todo
        var responseDeleteTodo = await _client.DeleteAsync($"/api/todos?id={_taskId}");
        responseDeleteTodo.StatusCode.Should().Be(System.Net.HttpStatusCode.OK);
        
        var responseDeleteUser = await _client.DeleteAsync($"/api/users?id={_userId}");
        responseDeleteUser.StatusCode.Should().Be(System.Net.HttpStatusCode.OK);
    }

    private CreateUserResource CreateSampleUser()
    {
        return new CreateUserResource()
        {
            FirstName = "Alex",
            LastName = "Brown",
            DailyTaskLimit = 10
        };
    }

    private TodoRequest CreateSampleTodoRequest(long userId)
    {
        return new TodoRequest
        {
            Name = "Test Todo",
            Description = "Test Description",
            UserId = userId
        };
    }
}