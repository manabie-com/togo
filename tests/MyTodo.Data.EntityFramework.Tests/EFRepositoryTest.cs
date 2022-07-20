using MyTodo.Data.Entities;
using MyTodo.Infrastructure.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using Xunit;

namespace MyTodo.Data.EntityFramework.Tests
{
    public class EFRepositoryTest
    {
        private readonly MyTodoDbContext _context;
        private readonly IUnitOfWork _unitOfWork;

        public EFRepositoryTest()
        {
            _context = ContextFactory.Create();
            _context.Database.EnsureCreated();
            _unitOfWork = new EFUnitOfWork(_context);
        }

        [Fact]
        public void Constructor_Should_Success_When_Create_EFRepository()
        {
            EFRepository<TodoItem, int> repository = new EFRepository<TodoItem, int>(_context);
            Assert.NotNull(repository);
        }

        [Fact]
        public void Add_Should_Have_Record_When_Insert()
        {
            EFRepository<TodoItem, int> repos = new EFRepository<TodoItem, int>(_context);
            repos.Add(new TodoItem()
            {
                Id = 1,
                Title = "Task 1",
                Description = "Task 1",
                Priority = 1,
                Status = Enums.TodoItemStatus.New
            });
            _unitOfWork.Commit();

            TodoItem function = repos.FindById(1);
            Assert.NotNull(function);

        }


        [Fact]
        public void FindAll_Should_Return_All_Record_In_Table()
        {
            EFRepository<TodoItem, int> repos = new EFRepository<TodoItem, int>(_context);

            repos.Add(new TodoItem()
            {
                Id = 1,
                Title = "Task 1",
                Description = "Task 1",
                Priority = 1,
                Status = Enums.TodoItemStatus.New
            });
            repos.Add(new TodoItem()
            {
                Id = 2,
                Title = "Task 2",
                Description = "Task 2",
                Priority = 1,
                Status = Enums.TodoItemStatus.New
            });
            _unitOfWork.Commit();

            List<TodoItem> todos = repos.FindAll().ToList();
            Assert.Equal(2, todos.Count);

        }

        [Fact]
        public void FindByIdAsync_Should_Return_True_Record_In_Table()
        {
            EFRepository<TodoItem, int> repos = new EFRepository<TodoItem, int>(_context);
            repos.Add(new TodoItem()
            {
                Id = 1,
                Title = "Task 1",
                Description = "Task 1",
                Priority = 1,
                Status = Enums.TodoItemStatus.New
            });
            _unitOfWork.Commit();

            TodoItem function = repos.FindById(1);
            Assert.Equal(1, function.Id);

        }


        [Fact]
        public void Update_Should_Have_Change_Record()
        {

            EFRepository<TodoItem, int> repos = new EFRepository<TodoItem, int>(_context);
            repos.Add(new TodoItem()
            {
                Id = 1,
                Title = "Task 1",
                Description = "Task 1",
                Priority = 1,
                Status = Enums.TodoItemStatus.New
            });
            _unitOfWork.Commit();
            TodoItem updated = repos.FindById(1);
            updated.Title = "Task 1 (updated)";
            repos.Update(updated);
            _unitOfWork.Commit();

            TodoItem function = repos.FindById(1);
            Assert.Equal("Task 1 (updated)", function.Title);
        }

        [Fact]
        public void Remove_Should_Success_When_Pass_Valid_Id()
        {

            EFRepository<TodoItem, int> repos = new EFRepository<TodoItem, int>(_context);
            TodoItem todo = new TodoItem()
            {

                Id = 1,
                Title = "Task 1",
                Description = "Task 1",
                Priority = 1,
                Status = Enums.TodoItemStatus.New
            };
            repos.Add(todo);
            _unitOfWork.Commit();

            repos.Remove(todo);
            _unitOfWork.Commit();

            TodoItem findTodo = repos.FindById(1);
            Assert.Null(findTodo);
        }



        [Fact]
        public void FindSingle_Should_Return_One_Record_If_Condition_Is_Match()
        {
            EFRepository<TodoItem, int> repos = new EFRepository<TodoItem, int>(_context);

            TodoItem todo = new TodoItem()
            {

                Id = 1,
                Title = "Task 1",
                Description = "Task 1",
                Priority = 1,
                Status = Enums.TodoItemStatus.New
            };
            repos.Add(todo);
            _unitOfWork.Commit();

            TodoItem result = repos.FindSingle(x => x.Title == "Task 1");
            Assert.NotNull(result);
        }


    }

}
