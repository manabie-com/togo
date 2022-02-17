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
        public DbSet<User> Users { get; set; }
        public DbSet<UserSettings> UserSettings { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<UserSettings>().HasKey(settings => settings.UserId);

            modelBuilder.Entity<Task>()
                .HasOne<User>()
                .WithMany(user => user.Tasks)
                .HasForeignKey(task => task.UserId);

            modelBuilder.Entity<User>()
                .HasOne(user => user.Settings)
                .WithOne(settings => settings.User)
                .HasForeignKey<UserSettings>(settings => settings.UserId);

            modelBuilder.Entity<User>()
                .HasData(new User { Id = "firstuser", FirstName = "Ha", LastName = "Nguyen" });
            modelBuilder.Entity<User>()
                .HasData(new User { Id = "seconduser", FirstName = "Ha", LastName = "Thanh Nguyen" });

            modelBuilder.Entity<UserSettings>()
                .HasData(new UserSettings { UserId = "firstuser", MaxTasksPerDay = 5 });
            modelBuilder.Entity<UserSettings>()
                .HasData(new UserSettings { UserId = "seconduser", MaxTasksPerDay = 15 });
        }
    }
}
