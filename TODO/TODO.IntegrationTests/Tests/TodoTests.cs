using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using TODO.Api;
using TODO.Repositories.Models.RequestModels;
using Xunit;
using System.Text.Json;
using System.Net;
using TODO.Repositories.Data.DBModels;

namespace TODO.IntegrationTests.Tests
{
    public class TodoTests : IClassFixture<CustomWebApplicationFactory<Startup>>
    {
        private readonly HttpClient _httpClient;
        private readonly CustomWebApplicationFactory<Startup> _factory;

        public TodoTests(CustomWebApplicationFactory<Startup> factory)
        {
            _factory = factory ?? throw new ArgumentNullException(nameof(factory));
            _httpClient = _factory.CreateClient();
        }

        private CreateTodoRequest GetTodoRequest(int userId)
        {
            return new CreateTodoRequest
            {
                UserId = userId,
                StatusId = 0,
                TodoName = "Test todo name",
                TodoDescription = "Test todo description",
                Priority = 0,
                DateCreated = DateTime.UtcNow,
                DateModified = null
            };
        }

        [Fact]
        public async Task CreateTodo_Returns201()
        {
            // Arrange
            var userId = 1;
            var content = new StringContent(JsonSerializer.Serialize(GetTodoRequest(userId)), Encoding.UTF8, "application/json");
            var serializeOptions = new JsonSerializerOptions
            {
                PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
            };

            // Act
            var response = await _httpClient.PostAsync("api/todos", content);
            var responseUsers = await _httpClient.GetAsync($"api/users?userId={userId}");

            var userListString = await responseUsers.Content.ReadAsStringAsync();
            var userList = JsonSerializer.Deserialize<List<User>>(userListString, serializeOptions);
            var user = userList.SingleOrDefault();
            var todo = user.Todos.Where(t => t.TodoId == 1).SingleOrDefault();

            // Assert
            Assert.Equal(HttpStatusCode.Created, response.StatusCode);
            Assert.Equal(HttpStatusCode.OK, responseUsers.StatusCode);
            Assert.Equal("Test todo name", todo.TodoName);
            Assert.Equal(response.Headers.Location.AbsoluteUri, $"https://localhost:5001/api/todos/{todo.TodoId}");
        }

    }
}
