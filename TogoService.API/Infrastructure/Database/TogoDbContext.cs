using Microsoft.EntityFrameworkCore;
using TogoService.API.Model;

namespace TogoService.API.Infrastructure.Database
{
    public class TogoDbContext : DbContext
    {
        public TogoDbContext(DbContextOptions<TogoDbContext> options) : base(options) { }

        public TogoDbContext() { }

        public DbSet<User> Users { get; set; }
        public DbSet<TodoTask> TodoTasks { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            base.OnModelCreating(modelBuilder);

            // Model
            modelBuilder.Entity<TodoTask>()
                .HasOne(task => task.User)
                .WithMany(user => user.Tasks)
                .HasForeignKey(task => task.UserId);
        }
    }
}
