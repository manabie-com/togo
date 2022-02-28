using Microsoft.VisualStudio.TestTools.UnitTesting;
using Microsoft.Extensions.Logging;
using Moq;
using TODO.Api.Controllers;
using Microsoft.AspNetCore.Mvc;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc.Infrastructure;
using TODO.Repositories.Repositories;
using TODO.Repositories.Interfaces;
using TODO.Repositories.Models.RequestModels;
using System;
using TODO.Repositories.Data.DBModels;
using System.Collections.Generic;

namespace TODO.UnitTests
{
    [TestClass]
    public class TodoControllerTests
    {
        private TodosController _sut;
        private readonly ILogger<TodosController> _mockLogger;
        private readonly IUserRepository _userRepoMock = Mock.Of<IUserRepository>();
        private readonly ITodoRepository _todoRepoMock = Mock.Of<ITodoRepository>();

        private CreateTodoRequest _request;
        private User _user;

        //private readonly Mock<IUserRepository> _userRepoMock = new Mock<IUserRepository>();
        //private readonly ITodoRepository _todoRepoMock = Mock.Of<ITodoRepository>();

        public TodoControllerTests()
        {
            _mockLogger = Mock.Of<ILogger<TodosController>>();
            _sut = new TodosController(_mockLogger, _todoRepoMock, _userRepoMock);

            _request = new CreateTodoRequest
            {
                UserId = 1,
                StatusId = 1,
                Priority = 1,
                TodoName = "Test Todo",
                TodoDescription = "Test Description",
                DateCreated = DateTime.UtcNow,
                DateModified = null
            };

            _user = Mock.Of<User>();
            _user.UserId = 1;
            _user.UserTodoConfig = new UserTodoConfig
            {
                UserId = 1,
                DailyTaskLimit = 3
            };
        }

        [TestMethod]
        public async Task CreateTodo_IsSuccessful_WhenLimitIsNotReached()
        {
            // Arrange
            _user.Todos = new List<Todo>
            {
                new Todo { TodoId = 1, UserId = 1, StatusId = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow, DateModified = null },
                new Todo { TodoId = 2, UserId = 1, StatusId = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow, DateModified = null },
            };

            var userList = new List<User>();
            userList.Add(_user);

            Mock.Get(_userRepoMock).Setup(x => x.GetUsers(1)).ReturnsAsync(userList);
            Mock.Get(_todoRepoMock).Setup(x => x.CreateTodo(_request)).ReturnsAsync(new Todo { TodoId = 3 });

            // Act
            var result = (IStatusCodeActionResult)await _sut.CreateTodo(_request);

            // Assert
            Assert.AreEqual(201, result.StatusCode);
        }

        [TestMethod]
        public async Task CreateTodo_ReturnsBadRequest_WhenLimitIsReached()
        {
            // Arrange
            _user.Todos = new List<Todo>
            {
                new Todo { TodoId = 1, UserId = 1, StatusId = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow, DateModified = null },
                new Todo { TodoId = 2, UserId = 1, StatusId = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow, DateModified = null },
                new Todo { TodoId = 2, UserId = 1, StatusId = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow, DateModified = null },
            };

            var userList = new List<User>();
            userList.Add(_user);

            Mock.Get(_userRepoMock).Setup(x => x.GetUsers(1)).ReturnsAsync(userList);

            // Act
            var result = (IStatusCodeActionResult)await _sut.CreateTodo(_request);

            // Assert
            Assert.AreEqual(400, result.StatusCode);
        }


        [TestMethod]
        public async Task CreateTodo_ReturnsCreated_IfTodosAreGreaterThanLimitButDifferentDate()
        {
            // Arrange
            _user.Todos = new List<Todo>
            {
                new Todo { TodoId = 1, UserId = 1, StatusId = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow.AddDays(-1), DateModified = null },
                new Todo { TodoId = 2, UserId = 1, StatusId = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow, DateModified = null },
                new Todo { TodoId = 2, UserId = 1, StatusId = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow, DateModified = null },
            };

            var userList = new List<User>();
            userList.Add(_user);

            Mock.Get(_userRepoMock).Setup(x => x.GetUsers(1)).ReturnsAsync(userList);
            Mock.Get(_todoRepoMock).Setup(x => x.CreateTodo(_request)).ReturnsAsync(new Todo { TodoId = 3 });

            // Act
            var result = (IStatusCodeActionResult)await _sut.CreateTodo(_request);

            // Assert
            Assert.AreEqual(201, result.StatusCode);
        }
    }
}
