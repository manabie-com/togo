using Models;
using Moq;
using Repositories;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace UnitTests.Data
{
    public class Initializer
    {
        public void Initialize(Mock<DatabaseContext> mockContext)
        {
            CreateUser(mockContext);
        }

        private void CreateUser(Mock<DatabaseContext> mockContext)
        {
            var users = new List<Users>()
            {
                new Users()
                {
                    Id = "1",
                    Username = "Admin",
                    Password = BCrypt.Net.BCrypt.HashPassword("12345678", BCrypt.Net.BCrypt.GenerateSalt(10)),
                    TaskPerDay = 10
                },
                new Users()
                {
                    Id = "2",
                    Username = "TestUser",
                    Password = BCrypt.Net.BCrypt.HashPassword("12345678", BCrypt.Net.BCrypt.GenerateSalt(10)),
                    TaskPerDay = 2
                }
            };

            var mockUserDbSet = GetQueryableMockDbSet.GetQueryableMockDbSets(users);

            mockUserDbSet.Setup(d => d.Attach(It.IsAny<Users>())).Callback<Users>((s) =>
            {
                var itemUpdate = users.Find(x => x.Id == s.Id);
                itemUpdate.TaskPerDay = s.TaskPerDay;
            });

            mockUserDbSet.Setup(d => d.Remove(It.IsAny<Users>())).Callback<Users>((s) =>
            {
                var user = users.Find(x => x.Id == s.Id);
                users.Remove(user);
            });

            var tasks = new List<Tasks>()
            {
                new Tasks()
                {
                    ID = "1",
                    Content = "Task Content",
                    UserID = "2",
                    CreateAt = DateTime.Now
                },
                new Tasks()
                {
                    ID = "2",
                    Content = "Task Content",
                    UserID = "2",
                    CreateAt = DateTime.Now
                }
            };

            var mockTaskDbSet = GetQueryableMockDbSet.GetQueryableMockDbSets(tasks);

            mockContext.Setup(x => x.Users).Returns(mockUserDbSet.Object);
            mockContext.Setup(x => x.Tasks).Returns(mockTaskDbSet.Object);
        }
    }
}
