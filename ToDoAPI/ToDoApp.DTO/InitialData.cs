using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using ToDoApp.DTO.Entity;

namespace ToDoApp.DTO
{
    public static class InitialData
    {
        public static void CreateSampleData(this DatabaseContext context)
        {
            context.Database.EnsureDeleted();
            context.Users.Add(new User
            {
                UserId = 1,
                DailyLimit = 10
            });
            context.SaveChanges();

            context.Todos.AddRange(
                new ToDo
                {
                Id = 1,
                Title = "Item 1",
                Detail = "Detail 1",
                CreatedDate = DateTime.Now,
                UserId =1
                },
                new ToDo
                {
                    Id = 2,
                    Title = "Item 2",
                    Detail = "Detail 2",
                    CreatedDate = DateTime.Now,
                    UserId = 1
                },
                new ToDo
                {
                    Id = 3,
                    Title = "Item 3",
                    Detail = "Detail 3",
                    CreatedDate = DateTime.Now,
                    UserId = 1
                },
                new ToDo
                {
                    Id = 4,
                    Title = "Item 4",
                    Detail = "Detail 4",
                    CreatedDate = DateTime.Now,
                    UserId = 1
                }
            );
            context.SaveChanges();
        }

        public static void CreateMockSampleData(this DatabaseContext context)
        {

            context.Users.Add(new User
            {
                UserId = 1,
                DailyLimit = 10
            });
            context.SaveChanges();

            context.Todos.AddRange(
                new ToDo
                {
                    Id = 2,
                    Title = "Item 2",
                    Detail = "Detail 2",
                    CreatedDate = DateTime.Now,
                    UserId = 1
                },
                new ToDo
                {
                    Id = 3,
                    Title = "Item 3",
                    Detail = "Detail 3",
                    CreatedDate = DateTime.Now,
                    UserId = 1
                }, 
                new ToDo
                {
                    Id = 4,
                    Title = "Item 4",
                    Detail = "Detail 4",
                    CreatedDate = DateTime.Now,
                    UserId = 1
                }
            );
            context.SaveChanges();
        }
    }
}
