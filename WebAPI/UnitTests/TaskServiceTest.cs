using Models;
using Moq;
using Repositories.Infrastructure;
using Services.Implementations;
using Services.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Text;
using System.Threading.Tasks;
using Xunit;

namespace UnitTests
{
    public class TaskServiceTest : BaseService_Test, IDisposable
    {
        private ITaskService _taskService;
        private Mock<IRepository<Tasks>> _mockTaskRepository;
        private Mock<IRepository<Users>> _mockUserRepository;

        protected override void Setup()
        {
            _mockTaskRepository = new Mock<IRepository<Tasks>>();
            _mockUserRepository = new Mock<IRepository<Users>>();

            _mockTaskRepository.Setup(x => x.FindAll(It.IsAny<Expression<Func<Tasks, bool>>>()))
                               .Returns(new Func<Expression<Func<Tasks, bool>>, IQueryable<Tasks>>(
                                   (predicate) =>
                                   {
                                       return _mockContext.Object.Tasks.Where(predicate);
                                   }));

            _mockTaskRepository.Setup(x => x.FindSingle(It.IsAny<Expression<Func<Tasks, bool>>>()))
                               .Returns(new Func<Expression<Func<Tasks, bool>>, Tasks>(
                                   (predicate) =>
                                   {
                                       var task = _mockContext.Object.Tasks;

                                       return task.FirstOrDefault(predicate);
                                   }));

            _mockTaskRepository.Setup(x => x.Add(It.IsAny<Tasks>()));

            _mockTaskRepository.Setup(x => x.Count(It.IsAny<Expression<Func<Tasks, bool>>>()))
                               .Returns(new Func<Expression<Func<Tasks, bool>>, int>(
                                   (predicate) =>
                                   {
                                       return _mockContext.Object.Tasks.Count(predicate);
                                   }));

            _mockTaskRepository.Setup(x => x.DbSet).Returns(_mockContext.Object.Tasks);

            _mockUserRepository.Setup(x => x.FindSingle(It.IsAny<Expression<Func<Users, bool>>>()))
                              .Returns(new Func<Expression<Func<Users, bool>>, Users>(
                                  (predicate) =>
                                  {
                                      var user = _mockContext.Object.Users;

                                      return user.FirstOrDefault(predicate);
                                  }));

            _mockLoggerFactory.Setup(l => l.CreateLogger("TaskService")).Returns(_mockLogger.Object);

            _taskService = new TaskService(_mockTaskRepository.Object, _mockUnitOfWork.Object, _mockUserRepository.Object, _mockLoggerFactory.Object);
        }

        [Fact]
        public void Should_Create_Success()
        {
            var task = new Tasks()
            {
                Content = "Task Content",
            };

            var result = _taskService.Create(task, "1");

            Assert.Equal(1, result);
        }

        [Fact]
        public void Should_Create_Fail()
        {
            var task = new Tasks()
            {
                Content = "Task Content",
            };

            var result = _taskService.Create(task, "2");

            Assert.Equal(0, result);
        }

        [Fact]
        public void Should_Create_Exception()
        {
            var task = new Tasks()
            {
                Content = "Task Content",
            };

            var result = _taskService.Create(task, "5");

            Assert.Equal(-1, result);
        }

        [Fact]
        public void Should_GetTaskByUserId_HaveData()
        {
            var result = _taskService.GetTasksByUserId("2");

            Assert.True(result.Count > 0);
        }

        [Fact]
        public void Should_GetTaskByUserId_HaveNoData()
        {
            var result = _taskService.GetTasksByUserId("3");

            Assert.True(result.Count == 0);
        }

        public void Dispose()
        {
            GC.SuppressFinalize(this);
        }
    }
}
