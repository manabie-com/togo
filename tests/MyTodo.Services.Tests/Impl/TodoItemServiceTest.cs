using AutoMapper;
using Moq;
using MyTodo.Data.Entities;
using MyTodo.Data.Enums;
using MyTodo.Infrastructure.Interfaces;
using MyTodo.Services.Config.AutoMapper;
using MyTodo.Services.Impl;
using MyTodo.Services.ViewModels;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Text;
using Xunit;

namespace MyTodo.Services.Tests.Impl
{
    public class TodoItemServiceTest
    {
        private readonly Mock<IRepository<TodoItem, int>> _mockTodoItemRepository;
        private readonly Mock<IUnitOfWork> _mockUnitOfWork;

        public TodoItemServiceTest()
        {
            _mockTodoItemRepository = new Mock<IRepository<TodoItem, int>>();
            _mockUnitOfWork = new Mock<IUnitOfWork>();

        }
        [Fact]
        public void Add_ValidInput_SucessResult()
        {
            var mockMapper = new MapperConfiguration(cfg => cfg.AddProfile(new ViewToModelMappingProfile()));
            var mapper = mockMapper.CreateMapper();
            _mockTodoItemRepository.Setup(x => x.Add(It.IsAny<TodoItem>()));
            var todoItemService = new TodoItemService(_mockTodoItemRepository.Object, _mockUnitOfWork.Object, mapper);
            var result = todoItemService.Add(
                new TodoItemViewModel()
                {
                    Id = 0,
                    Title = "Task 1",
                    Description = "Task 1",
                    Priority = 1,
                    Status = TodoItemStatus.New
                });
            Assert.NotNull(result);
        }

        [Fact]
        public void GetById_ValidQuery_ResultSuccess()
        {
            var mockMapper = new MapperConfiguration(cfg => cfg.AddProfile(new ModelToViewMappingProfile()));
            var mapper = mockMapper.CreateMapper();

            _mockTodoItemRepository.Setup(x => x.FindById(It.IsAny<int>()))
                .Returns(new TodoItem
                {
                    Id = 1,
                    Title = "Task 1",
                    Description = "Task 1",
                    Priority = 1,
                    Status = TodoItemStatus.New
                });

            var todoItemService = new TodoItemService(_mockTodoItemRepository.Object, _mockUnitOfWork.Object, mapper);
            var result = todoItemService.GetById(1);

            Assert.Equal(1, result.Id);
        }
    }
}
