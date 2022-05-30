using Microsoft.AspNetCore.Mvc.Infrastructure;
using Microsoft.Extensions.Logging;
using MongoDB.Driver;
using Moq;
using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using ToDo.Api.Controllers;
using ToDo.Api.Domain.DBModels;
using ToDo.Api.Repositories;
using ToDo.Api.Requests;
using ToDo.Api.Validators;
using TODO.Repositories.Data.DBModels;
using Xunit;

namespace ToDo.UnitTest
{
    public class ToDosControllerTest
    {
        private TodosController _toDosController;
        private readonly ILogger<TodosController> _mockLogger;
        private readonly IMongoBaseRepository<User> _userRepoMock;
        private readonly IMongoBaseRepository<Todo> _toDoRepoMock;
        private readonly CreateToDoValidator _createToDoValidator;

        private CreateTodoRequest _request;
        private User _user;

        public ToDosControllerTest()
        {
            _mockLogger = new Mock<ILogger<TodosController>>().Object;
            _userRepoMock = new Mock<IMongoBaseRepository<User>>().Object;
            _toDoRepoMock = new Mock<IMongoBaseRepository<Todo>>().Object;
            _createToDoValidator = new CreateToDoValidator();
            _toDosController = new TodosController(_mockLogger, _toDoRepoMock, _userRepoMock, _createToDoValidator);
            _request = new CreateTodoRequest
            {
                UserId = Guid.NewGuid(),
                Status = 1,
                TodoName = "Test Todo",
                TodoDescription = "Test Description"
            };
        }
        [Fact]
        public async Task CreateTodo_IsSuccessful_WhenLimitIsNotReached()
        {
            // Arrange
            var todos = new List<Todo>
            {
                new Todo { UserId = Guid.NewGuid(), Status = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow},
                new Todo { UserId = Guid.NewGuid(), Status = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow},
            };
            _user = new User
            {
                FullName = "Nguyen Van A",
                DailyTaskLimit = 10,
                DateCreated = DateTime.UtcNow
            };

            Mock.Get(_userRepoMock).Setup(x => x.GetByIdAsync(It.IsAny<Guid>(), It.IsAny<CancellationToken>()))
                .ReturnsAsync(_user);
            Mock.Get(_toDoRepoMock).Setup(x => x.FindAllWithFilter(It.IsAny<FilterDefinition<Todo>>()))
                .ReturnsAsync(todos);

            // Act
            var result = (IStatusCodeActionResult)await _toDosController.CreateTodo(_request);

            // Assert
            Assert.Equal(200, result.StatusCode);
        }

        [Fact]
        public async Task CreateTodo_ReturnsBadRequest_WhenLimitIsReached()
        {
            // Arrange
            var todos = new List<Todo>
            {
                new Todo { UserId = Guid.NewGuid(), Status = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow},
                new Todo { UserId = Guid.NewGuid(), Status = 1, TodoName = "Test", TodoDescription = "Test", DateCreated = DateTime.UtcNow},
            };
            _user = new User
            {
                FullName = "Nguyen Van A",
                DailyTaskLimit = 2,
                DateCreated = DateTime.UtcNow
            };

            Mock.Get(_userRepoMock).Setup(x => x.GetByIdAsync(It.IsAny<Guid>(), It.IsAny<CancellationToken>()))
                .ReturnsAsync(_user);
            Mock.Get(_toDoRepoMock).Setup(x => x.FindAllWithFilter(It.IsAny<FilterDefinition<Todo>>()))
                .ReturnsAsync(todos);

            // Act
            var result = (IStatusCodeActionResult)await _toDosController.CreateTodo(_request);

            // Assert
            Assert.Equal(400, result.StatusCode);
        }

        [Fact]
        public async Task CreateTodo_ReturnsBadRequest_WhenValidator()
        {
            // Arrange
            var request = new CreateTodoRequest
            {
                UserId = Guid.NewGuid(),
                Status = 1,
                TodoName = "Test Todo Test Todo Test Todo Test Todo Test Todo" +
                "Test Todo Test Todo Test Todo Test Todo Test Todo Test Todo",
                TodoDescription = "Test Description"
            };

            // Act
            var result = (IStatusCodeActionResult)await _toDosController.CreateTodo(request);

            // Assert
            Assert.Equal(400, result.StatusCode);
        }
    }
}