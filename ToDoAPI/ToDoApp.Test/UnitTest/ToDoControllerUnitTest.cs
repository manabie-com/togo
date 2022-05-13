using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;
using ToDoApp.API.Controllers;
using ToDoApp.DTO.Entity;
using Xunit;

namespace ToDoApp.Test.UnitTest
{
    public class ToDoControllerUnitTest
    {
        private readonly ToDoController _controller;
        public ToDoControllerUnitTest()
        {
            _controller = new ToDoController(MockDatabaseContext.GetDatabaseContext());
        }

        [Fact]
        public async Task ToDoControllerUnitTest_GetToDoList()
        {
            var todos = await _controller.GetToDoList();
            Assert.NotNull(todos);
        }

        [Theory]
        [InlineData(2)]
        public async Task ToDoControllerUnitTest_GetToDoById(int value)
        {
            var todos = await _controller.GetToDoById(value);

            Assert.NotNull(todos);
            Assert.Equal(value, todos.Value.Id);

        }
        [Theory]
        [InlineData(11, "Title 11", "Detail 11", 1)]
        public async Task ToDoControllerUnitTest_Post(int @id, string title, string detail, int userId)
        {
            var newItem = new ToDo
            {
                Id = @id,
                Title = title,
                Detail = detail,
                UserId = userId
            };

            var response = await _controller.Post(newItem);
            
            var newTodos = await _controller.GetToDoById(@id);

            Assert.Equal(newItem.Id, newTodos.Value.Id);
            Assert.Equal(newItem.Title, newTodos.Value.Title);
            Assert.Equal(newItem.Detail, newTodos.Value.Detail);

        }
    }
}
