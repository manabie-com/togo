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
        DbSet<User> Users { get; set; }
        DbSet<User> Todos { get; set; }
        DbSet<UserTodoConfig> UserTodoConfig { get; set; }
        DbSet<TodoStatus> TodoStatus { get; set; }

        public TodoContext(DbContextOptions<TodoContext> options) : base(options)
        {
            
        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
        }
    }
}
