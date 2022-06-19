using AutoMapper;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;
using Moq;
using System;
using System.Collections.Generic;
using Todo.Application.Extensions;
using Todo.Application.Interfaces;
using Todo.Application.Services;
using Todo.Domain.Entities;
using Todo.Infrastructure;
using Xunit;

namespace Todo.Application.Test
{
    public class UserTaskServiceTest
    {
        private UserTaskService mockTaskService;
        private Mock<DbSet<UserTask>> mockSet;
        private Mock<IApplicationDbContext> mockDbContext;
        private Mock<IMapper> mapper;
        private DateTime currentDate;
        public UserTaskServiceTest()
        {
            mockSet = new Mock<DbSet<UserTask>>();
            mockDbContext = new Mock<IApplicationDbContext>();
            mockDbContext.Setup(m => m.UserTasks).Returns(mockSet.Object);

            mapper = new Mock<IMapper>();

            ResolverFactory.ServiceCollection = new ServiceCollection();
            ResolverFactory.ServiceCollection.AddTransient(scp => mapper.Object);

            mockTaskService = new UserTaskService(mapper.Object, mockDbContext.Object);

            currentDate = DateTime.UtcNow;
        }
        private void FakeDataContext(ApplicationDbContext context)
        {
            context.Users.Add(new User { Id = "001", Name = "UserTest", CreatedBy = "System", CreatedAt = currentDate });
            context.UserTasks.AddRange(GetListTask(currentDate));
            context.SaveChanges();
        }
        private ApplicationDbContext CreateDbContext()
        {
            var options = new DbContextOptionsBuilder<ApplicationDbContext>()
                .UseInMemoryDatabase(Guid.NewGuid().ToString("N")).Options;
            var dbContext = new ApplicationDbContext(options);
            return dbContext;
        }
        [Fact]
        public void UserId_Is_Exists()
        {
            //Arrange
            var context = CreateDbContext();

            FakeDataContext(context);

            mockTaskService = new UserTaskService(mapper.Object, context);

            var user = mockTaskService.GetUserById("001");

            Assert.NotNull(user);
            Assert.Equal("UserTest", user.Name);

            //Clean up
            context.Dispose();
        }
        [Fact]
        public void UserId_Is_Not_Exists()
        {
            //Arrange
            var context = CreateDbContext();

            FakeDataContext(context);

            mockTaskService = new UserTaskService(mapper.Object, context);

            Assert.Throws<ArgumentException>(() => mockTaskService.GetUserById("002"));

            context.Dispose();
        }

        [Fact]
        public void TotalTask_Is_Valid()
        {
            //Arrange
            var context = CreateDbContext();

            FakeDataContext(context);

            mockTaskService = new UserTaskService(mapper.Object, context);

            int total = mockTaskService.TotalCurrentTask(currentDate.Date, currentDate.Date.AddDays(1).AddTicks(-1L), "System");

            Assert.Equal(3, total);
            context.Dispose();
        }
        [Fact]
        public void Limit_task_is_larger_total_task()
        {
            bool status = mockTaskService.IsSmallerLimitTask(20, 19);
            Assert.True(status);
        }
        [Fact]
        public void Limit_task_is_smaller_total_task()
        {
            bool status = mockTaskService.IsSmallerLimitTask(15, 15);
            Assert.False(status);
        }
        [Fact]
        public void Create_a_task_via_context()
        {
            var task = new UserTask()
            {
                Title = "This is test",
                Description = "This is description test",
                Type = 0,
                Priority = 0,
            };
            mockTaskService.CreateNewTask(task);

            mockSet.Verify(m => m.Add(It.IsAny<UserTask>()), Times.Once());

            mockDbContext.Verify(m => m.SaveChanges(), Times.Once());
        }
        [Fact]
        public void CurrentDateRangeIsValid()
        {
            var date = new DateTime(2022, 06, 17);
            var dateRange = mockTaskService.GetCurrentDateRange(date);

            Assert.Equal(new DateTime(2022, 06, 17, 0, 0, 0), dateRange.StartDate);
            Assert.Equal((new DateTime(2022, 06, 17, 23, 59, 59)).AddSeconds(1).AddTicks(-1L), dateRange.EndDate);
        }
        private List<UserTask> GetListTask(DateTime createdDate)
        {
            var tasks = new List<UserTask> {
                    new UserTask()
                    {
                        Title = "This is test 01",
                        Description = "This is description test 01",
                        Type = 0,
                        Priority = 0,
                        CreatedAt = createdDate,
                        CreatedBy ="System",
                    },
                    new UserTask()
                    {
                        Title = "This is test 02",
                        Description = "This is description test 02",
                        Type = 0,
                        Priority = 0,
                        CreatedAt = createdDate,
                        ModifiedAt = DateTime.Now,
                        CreatedBy ="System",
                    },
                    new UserTask()
                    {
                        Title = "This is test 03",
                        Description = "This is description test 03",
                        Type = 0,
                        Priority = 0,
                        CreatedAt = createdDate,
                        CreatedBy ="System",
                    }
                    };
            return tasks;
        }
    }
}