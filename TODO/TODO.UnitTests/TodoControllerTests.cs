using Microsoft.VisualStudio.TestTools.UnitTesting;
using Microsoft.Extensions.Logging;
using Moq;
using TODO.Api.Controllers;
using Microsoft.AspNetCore.Mvc;
using System.Threading.Tasks;

namespace TODO.UnitTests
{
    [TestClass]
    public class TodoControllerTests
    {
        [TestMethod]
        public async Task TestRun()
        {
            // Arrange 
            var logger = Mock.Of<ILogger<WeatherForecastController>>();
            var weatherForecastController = new WeatherForecastController(logger);

            // Act
            var result = await weatherForecastController.GetSomething();

            // Assert
            Assert.IsInstanceOfType(result, typeof(OkResult));
        }
    }
}
