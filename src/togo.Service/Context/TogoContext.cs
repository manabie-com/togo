using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;

namespace togo.Service.Context
{
    public class TogoContext : DbContext
    {
        public TogoContext(DbContextOptions options) : base(options)
        {

        }

        public DbSet<Task> Tasks { get; set; }
        public DbSet<User> Users { get; set; }

        protected override void OnModelCreating(ModelBuilder builder)
        {
            builder.Entity<Task>(entity =>
            {
                entity.HasKey(e => e.Id);

                entity.Property(e => e.Id)
                      .IsRequired()
                      .HasMaxLength(256);
                entity.Property(e => e.Content)
                      .IsRequired();
            });

            builder.Entity<Task>()
                .HasOne(x => x.User)
                .WithMany(x => x.Tasks)
                .HasForeignKey(x => x.UserId);

            builder.Entity<User>(entity =>
            {
                entity.HasKey(e => e.Id);

                entity.Property(e => e.Id)
                      .IsRequired()
                      .HasMaxLength(256);

                entity.Property(e => e.Password)
                      .IsRequired()
                      .HasMaxLength(256);

                entity.Property(e => e.MaxTodo)
                      .IsRequired()
                      .HasDefaultValue(5);
            });

            base.OnModelCreating(builder);
        }
    }

    public class Task
    {
        public string Id { get; set; }
        public string UserId { get; set; }
        public string Content { get; set; }
        public DateTime CreatedDate { get; set; }

        public User User { get; set; }
    }

    public class User
    {
        public string Id { get; set; }
        public string Password { get; set; }
        public int MaxTodo { get; set; }

        public List<Task> Tasks { get; set; }
    }
}
