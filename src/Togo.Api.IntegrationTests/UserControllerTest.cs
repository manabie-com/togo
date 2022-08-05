using System;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc.Testing;
using Newtonsoft.Json;
using Togo.Infrastructure.Identities.Dtos;
using Xunit;

namespace Togo.Api.IntegrationTests;

public class UserControllerTest : IClassFixture<WebApplicationFactory<Program>>
{
    private readonly HttpClient _http;
    
    public UserControllerTest(WebApplicationFactory<Program> factory)
    {
        _http = factory.CreateDefaultClient();
    }

    [Fact]
    public async Task CreateAsync_WhenAllCorrect_ShouldSuccess()
    {
        var request = new HttpRequestMessage(HttpMethod.Post, "api/user");
        var createUserDto = new CreateUserDto
        {
            UserName = Guid.NewGuid().ToString(),
            Password = "Abcd@1234",
            MaxTasksPerDay = 10
        };

        request.Content = new StringContent(
            JsonConvert.SerializeObject(createUserDto), 
            Encoding.UTF8, 
            "application/json");

        var response = await _http.SendAsync(request);
        Assert.True(response.IsSuccessStatusCode);
    }
}
