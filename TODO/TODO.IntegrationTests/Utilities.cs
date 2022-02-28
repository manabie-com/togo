using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using TODO.Repositories.Data;
using TODO.Repositories.Data.DBModels;

namespace TODO.IntegrationTests
{
    public static class Utilities
    {
        public static async Task InitializeDbForTests(TodoContext db)
        {
            await db.User.AddRangeAsync(GetSeedingUsers());
            await db.UserTodoConfig.AddRangeAsync(GetSeedingUserTodoConfigs());
            await db.TodoStatus.AddRangeAsync(GetSeedingTodoStatus());
            await db.SaveChangesAsync();
        }

        public static async Task ReinitializeDbForTests(TodoContext db)
        {
            db.UserTodoConfig.RemoveRange(db.UserTodoConfig);
            db.User.RemoveRange(db.User);
            db.TodoStatus.RemoveRange(db.TodoStatus);
            db.Todo.RemoveRange(db.Todo);
            await db.SaveChangesAsync();
        }

        public static List<User> GetSeedingUsers()
        {
            return new List<User>
            {
                new User { UserId = 1, LastName = "Michael", FirstName = "Jordan" },
                new User { UserId = 2, LastName = "Barkley", FirstName = "Charles" },
            };
        }

        public static List<UserTodoConfig> GetSeedingUserTodoConfigs()
        {
            return new List<UserTodoConfig>
            {
                new UserTodoConfig { UserId = 1, DailyTaskLimit = 10 },
                new UserTodoConfig { UserId = 2, DailyTaskLimit = 5 },
            };
        }

        public static List<TodoStatus> GetSeedingTodoStatus()
        {
            return new List<TodoStatus>
            {
                new TodoStatus { TodoStatusId = 0, StatusName = "TO DO" },
                new TodoStatus { TodoStatusId = 1, StatusName = "DONE" },
                new TodoStatus { TodoStatusId = 2, StatusName = "IN PROGRESS" },
            };
        }
    }
}
