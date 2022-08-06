using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc.Testing;
using Togo.Infrastructure.Identities.Dtos;
using Xunit;

namespace Togo.Api.IntegrationTests;

public class UserControllerTest : BaseIntegrationTestClass
{
    public UserControllerTest(WebApplicationFactory<Program> factory) : base(factory)
    {
    }

    [Fact]
    public async Task CreateAsync_WhenAllCorrect_ShouldSuccess()
    {
        var createUserInput = new CreateUserDto
        {
            UserName = Guid.NewGuid().ToString(),
            Password = "Abcd@1234",
            MaxTasksPerDay = 10
        };

        var response = await PostAsync("api/user", createUserInput);
        
        Assert.True(response.IsSuccessStatusCode);
    }
}
