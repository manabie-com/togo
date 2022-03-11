using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using TODO.Repositories.Data.DBModels;


namespace TODO.Repositories.Data
{
    public class TodoContext : DbContext
    {
        public DbSet<User> User { get; set; }
        public DbSet<Todo> Todo { get; set; }
        public DbSet<UserTodoConfig> UserTodoConfig { get; set; }
        public DbSet<TodoStatus> TodoStatus { get; set; }

        public TodoContext(DbContextOptions<TodoContext> options) : base(options)
        {
            
        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<User>()
                .HasData(new User { UserId = 1, LastName = "Jordan", FirstName = "Michael" }, new { UserId = 2, LastName = "Thomas", FirstName = "Isiah"});

            modelBuilder.Entity<UserTodoConfig>()
                .HasData(new UserTodoConfig { UserId = 1, DailyTaskLimit = 10 }, new UserTodoConfig { UserId = 2, DailyTaskLimit = 5 });

            modelBuilder.Entity<TodoStatus>()
                .Property(o => o.TodoStatusId)
                .ValueGeneratedNever();

            modelBuilder.Entity<TodoStatus>()
                .HasData(new TodoStatus { TodoStatusId = 0, StatusName = "TO DO" }, new TodoStatus { TodoStatusId = 1, StatusName = "DONE" }, new TodoStatus { TodoStatusId = 2, StatusName = "IN PROGRESS" });

            base.OnModelCreating(modelBuilder);
        }
    }
}
