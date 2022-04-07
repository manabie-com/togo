using System;
using System.Net;
using System.Threading.Tasks;

using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.AspNetCore.TestHost;

using Xunit;

namespace Manabie.Api.IntegrationTests.SUT
{
    public class HeartBeatTest
    {
        [Fact]
        public async Task TestHeath()
        {
            // Arrange
            var factoty = new WebApplicationFactory<Program>();
            var client = factoty.CreateDefaultClient();

            //Act
            var response = await client.GetAsync("/");

            // Assert
            Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        }
    }
}

