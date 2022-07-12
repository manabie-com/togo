using Microsoft.AspNetCore.Mvc;
using Moq;
using MyTodo.BackendApi.Controllers;
using MyTodo.Services.Interfaces;
using MyTodo.Services.ViewModels;
using System;
using System.Collections.Generic;
using System.Text;
using Xunit;

namespace MyTodo.BackendApi.Tests.Controllers
{
    public class TodoItemsControllerTests
    {
        [Fact]
        public void Get_ValidRequest_OkResult()
        {
            var mockTodoItems = new Mock<ITodoItemService>();
            var mockAssignmentService = new Mock<IAssignmentService>();
            mockTodoItems.Setup(x => x.GetAll()).Returns(new List<TodoItemViewModel>()
            {
                new TodoItemViewModel(){Id = 1, Title="Task 1", Description="Task 1", Priority=1, Status=Data.Enums.TodoItemStatus.New},
                new TodoItemViewModel(){Id = 1, Title="Task 2", Description="Task 2", Priority=1, Status=Data.Enums.TodoItemStatus.New},
            });
            var controller = new TodoItemsController(mockTodoItems.Object, mockAssignmentService.Object);
            var result = controller.GetAll();
            Assert.IsType<OkObjectResult>(result);
            Assert.Equal(200, (result as OkObjectResult).StatusCode);
        }

        [Fact]
        public void Get_ServiceException_BadRequestResult()
        {
            var mockTodoItems = new Mock<ITodoItemService>();
            var mockAssignmentService = new Mock<IAssignmentService>();

            mockTodoItems.Setup(x => x.GetAll()).Throws<Exception>();
            var controller = new TodoItemsController(mockTodoItems.Object, mockAssignmentService.Object);
            Assert.ThrowsAny<Exception>(() => { controller.GetAll(); });
        }
    }
}
