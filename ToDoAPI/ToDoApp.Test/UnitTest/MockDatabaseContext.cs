using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using ToDoApp.DTO;
using ToDoApp.DTO.Entity;

namespace ToDoApp.Test.UnitTest
{
    public static class MockDatabaseContext
    {
        public static readonly List<ToDo> _toDos = new List<ToDo>()
        {
            new ToDo()
                {
                    Id = 7,
                    Title = "Title 7",
                    Detail = "Detail 7",
                    CreatedDate = DateTime.Now,
                    UserId = 1
                },
                new ToDo()
                {
                    Id = 8,
                    Title = "Title 8",
                    Detail = "Detail 8",
                    CreatedDate = DateTime.Now,
                    UserId = 1
                },
                new ToDo()
                {
                    Id = 9,
                    Title = "Title 9",
                    Detail = "Detail 9",
                    CreatedDate = DateTime.Now,
                    UserId = 1
                }
        };

        public static DatabaseContext GetDatabaseContext()
        {
            var options = new DbContextOptionsBuilder<DatabaseContext>()
                    .UseInMemoryDatabase(Guid.NewGuid().ToString())
                    .Options;
            var databaseContext = new DatabaseContext(options);

            databaseContext.Database.EnsureCreated();
            databaseContext.Database.EnsureDeleted();
            databaseContext.CreateMockSampleData();
            if (databaseContext.Users.Count() <= 0)
            {
                databaseContext.Users.Add(new User()
                {
                    DailyLimit = 10,
                    UserId = 1
                });
                databaseContext.SaveChanges();
            }
            if (databaseContext.Todos.Count() <= 0)
            {
                databaseContext.Todos.AddRange(_toDos);
                databaseContext.SaveChanges();
            }
            return databaseContext;
        }
    }
}
