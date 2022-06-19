using AutoMapper;
using Microsoft.Extensions.DependencyInjection;
using Moq;
using System;
using Todo.Application.Extensions;
using Todo.Application.Interfaces;
using Todo.Application.Services;
using Xunit;

namespace Testing
{
    public class UserTaskServiceTest
    {
        public UserTaskService mockTaskService;
        public UserTaskServiceTest()
        {
            var mapper = new Mock<IMapper>();
            var dbContext = new Mock<IApplicationDbContext>();
            ResolverFactory.ServiceCollection = new ServiceCollection();
            ResolverFactory.ServiceCollection.AddTransient(scp => mapper.Object);
            mockTaskService = new UserTaskService(mapper.Object, dbContext.Object);
        }
        [Fact]
        public void CanCreateTask()
        {

        }
        public void CanGetUserById()
        {

        }
        public void CanCountTasksByUserId()
        {

        }
        public void CannotCountTasksByUserId()
        {

        }
        [Fact]
        public void CurrentDateRangeIsValid()
        {
            var date = new DateTime(2022, 06, 17);
            var dateRange = mockTaskService.GetCurrentDateRange(date);

            Assert.Equal(new DateTime(2022, 06, 17, 0, 0, 0), dateRange.StartDate);
            Assert.Equal((new DateTime(2022, 06, 17, 23, 59, 59)).AddSeconds(1).AddTicks(-1L), dateRange.EndDate);
        }
    }
}