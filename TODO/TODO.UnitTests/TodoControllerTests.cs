using Microsoft.VisualStudio.TestTools.UnitTesting;
using Microsoft.Extensions.Logging;
using Moq;
using TODO.Api.Controllers;
using Microsoft.AspNetCore.Mvc;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc.Infrastructure;

namespace TODO.UnitTests
{
    [TestClass]
    public class TodoControllerTests
    {
        //[TestMethod]
        //public async Task TestRun()
        //{
        //    // Arrange 
        //    var logger = Mock.Of<ILogger<WeatherForecastController>>();
        //    var weatherForecastController = new WeatherForecastController(logger);

        //    // Act
        //    var result = await weatherForecastController.GetSomething();

        //    // Assert
        //    Assert.IsInstanceOfType(result, typeof(OkResult));
        //}

        [TestMethod]
        public async Task TestRun_NotImplemented()
        {
            // Arrange 
            var logger = Mock.Of<ILogger<TodosController>>();
            var todosController = new TodosController(logger);

            // Act
            var result = (IStatusCodeActionResult)await todosController.CreateTodo();

            // Assert
            Assert.AreEqual(result.StatusCode, 500);
        }
    }
}
