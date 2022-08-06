using System.Net.Http;
using System.Net.Http.Headers;
using System.Text;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Mvc.Testing;
using Newtonsoft.Json;
using Togo.Infrastructure.Identities.Dtos;
using Xunit;

namespace Togo.Api.IntegrationTests;

public abstract class BaseIntegrationTestClass : IClassFixture<WebApplicationFactory<Program>>
{
    protected const string CreateUserPath = "api/user";
    protected const string LoginPath = "api/user/login";
    protected const string AdminUserName = "admin";
    protected const string AdminPassword = "Abcd@1234";
    protected readonly HttpClient HttpClient;

    protected BaseIntegrationTestClass(WebApplicationFactory<Program> webApplicationFactory)
    {
        HttpClient = webApplicationFactory.CreateDefaultClient();
    }

    protected async Task<HttpResponseMessage> PostAsync<T>(string path, T body)
    {
        var request = new HttpRequestMessage(HttpMethod.Post, path);
        
        request.Content = new StringContent(
            JsonConvert.SerializeObject(body), 
            Encoding.UTF8,
            "application/json");
        
        return await HttpClient.SendAsync(request);
    }

    public async Task LoginAs(string username, string password)
    {
        var loginRequest = new LoginDto
        {
            UserName = username,
            Password = password
        };

        var response = await PostAsync(LoginPath, loginRequest);
        response.EnsureSuccessStatusCode();

        var responseString = await response.Content.ReadAsStringAsync();
        var responseObject = JsonConvert.DeserializeObject<LoginResponseDto>(responseString);

        HttpClient.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue(
            JwtBearerDefaults.AuthenticationScheme, responseObject.AccessToken);
    }
}
