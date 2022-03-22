using Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Repositories.Data
{
    public static class Seeder
    {
        public static void Seed(this DatabaseContext context)
        {
            context.Database.EnsureCreated();
            
            context.Users.Add(new Users()
            {
                Id = "1",
                Username = "TestUser1",
                Password = BCrypt.Net.BCrypt.HashPassword("123456789"),
                TaskPerDay = 10
            });

            context.Users.Add(new Users()
            {
                Id = "2",
                Username = "TestUser2",
                Password = BCrypt.Net.BCrypt.HashPassword("123456789"),
                TaskPerDay = 2
            });
            
            context.Tasks.Add(new Tasks()
            {
                ID = "1",
                Content = "Task Content",
                UserID = "2",
                CreateAt = DateTime.Now
            });

            context.Tasks.Add(new Tasks()
            {
                ID = "2",
                Content = "Task Content",
                UserID = "2",
                CreateAt = DateTime.Now
            });

            context.SaveChanges();
        }
    }
}
