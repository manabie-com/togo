using Microsoft.EntityFrameworkCore;
using ToDoBackend.Models;

namespace ToDoBackend
{
    public class ToDoContext : DbContext
    {
        public ToDoContext(DbContextOptions<ToDoContext> options) : base(options)
        {
        }

        public DbSet<Task> Tasks { get; set; }
    }
}
