using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.VisualStudio.TestPlatform.TestHost;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using ToDoApp.API;
using ToDoApp.API.Controllers;
using ToDoApp.DTO.Entity;
using ToDoApp.Test.IntegrationTest;
using Xunit;

namespace ToDoApp.Test
{
    public class ToDoControllerIntergrationTest : IClassFixture<WebApplicationFactory<Startup>>
    {
        private readonly HttpClient _client;
        private readonly WebApplicationFactory<Startup> _applicationFactory;

        public ToDoControllerIntergrationTest()
        {
            var webAppFactory = new WebApplicationFactory<Startup>();
            _client = webAppFactory.CreateClient();  


        }
        [Fact]
        public async Task ToDoController_Get_Success()
        {
            var expected = new ToDo
            {
               Id = 1,
               Title = "Item 1",
               Detail = "Detail 1"
            };

            var response = await _client.GetAsync("/api/ToDo/");

            response.EnsureSuccessStatusCode();
            var objectResponse = JsonConvert.DeserializeObject<List<ToDo>>(response.Content.ReadAsStringAsync().Result).FirstOrDefault();

            Assert.NotNull(objectResponse);
            Assert.Equal(expected.Id, objectResponse.Id);
            Assert.Equal(expected.Title, objectResponse.Title);
            Assert.Equal(expected.Detail, objectResponse.Detail);
        }

        [Fact]
        public async Task ToDoController_Post_Success()
        {
            var expected = new ToDo
            {
                Id = 5,
                Title = "Item 5",
                Detail = "Detail 5",
                UserId = 1
            };

            var postRequest = new HttpRequestMessage(HttpMethod.Post, "/api/ToDo");
            postRequest.Content = new StringContent(JsonConvert.SerializeObject(expected), Encoding.UTF8, "application/json"); 
            var postResponse = await _client.SendAsync(postRequest);

            postResponse.EnsureSuccessStatusCode();

            var response = await _client.GetAsync("/api/ToDo/5");

            response.EnsureSuccessStatusCode();
            var objectResponse = JsonConvert.DeserializeObject<ToDo>(response.Content.ReadAsStringAsync().Result);

            Assert.Equal(HttpStatusCode.OK, postResponse.StatusCode);
            Assert.Equal(expected.Id, objectResponse.Id);
            Assert.Equal(expected.Title, objectResponse.Title);
            Assert.Equal(expected.Detail, objectResponse.Detail);
        }

        [Fact]
        public async Task ToDoController_Put_Success()
        {
            var expected = new ToDo
            {
                Id = 1,
                Title = "New Item 1",
                Detail = "New Detail 1"
            };

            var postRequest = new HttpRequestMessage(HttpMethod.Put, "/api/ToDo/1");
            postRequest.Content = new StringContent(JsonConvert.SerializeObject(expected), Encoding.UTF8, "application/json");
            var postResponse = await _client.SendAsync(postRequest);

            postResponse.EnsureSuccessStatusCode();

            var response = await _client.GetAsync("/api/ToDo/1");

            response.EnsureSuccessStatusCode();
            var objectResponse = JsonConvert.DeserializeObject<ToDo>(response.Content.ReadAsStringAsync().Result);

            Assert.Equal(HttpStatusCode.OK, postResponse.StatusCode);
            Assert.Equal(expected.Id, objectResponse.Id);
            Assert.Equal(expected.Title, objectResponse.Title);
            Assert.Equal(expected.Detail, objectResponse.Detail);
        }

        [Fact]
        public async Task ToDoController_Delete_Success()
        {
            var expected = new ToDo
            {
                Id = 4,
                Title = "New Item 4",
                Detail = "Detail 4"
            };

            var postResponse = await _client.DeleteAsync("/api/ToDo/4");

            var response = await _client.GetAsync("/api/ToDo/4");

            var objectResponse = JsonConvert.DeserializeObject<ToDo>(response.Content.ReadAsStringAsync().Result);

            Assert.Equal(HttpStatusCode.OK, postResponse.StatusCode);
            Assert.Equal(HttpStatusCode.NotFound, response.StatusCode);
        }
    }
}
