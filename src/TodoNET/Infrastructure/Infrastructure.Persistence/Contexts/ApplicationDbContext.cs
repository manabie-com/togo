using Application.Interfaces;
using Domain.Common;
using Domain.Entities;
using Microsoft.EntityFrameworkCore;
using System;
using System.Threading;
using System.Threading.Tasks;
using TaskEntity = Domain.Entities.Task;

namespace Infrastructure.Persistence.Contexts
{
    public class ApplicationDbContext : DbContext
    {
        private readonly IAuthenticatedUserService _authenticatedUser;

        public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options, IAuthenticatedUserService authenticatedUser) : base(options)
        {
            ChangeTracker.QueryTrackingBehavior = QueryTrackingBehavior.NoTracking;
            _authenticatedUser = authenticatedUser;
        }
        public DbSet<User> Users { get; set; }
        public DbSet<TaskEntity> Tasks { get; set; }

        public override Task<int> SaveChangesAsync(CancellationToken cancellationToken = new CancellationToken())
        {
            foreach (var entry in ChangeTracker.Entries<AuditableBaseEntity>())
            {
                switch (entry.State)
                {
                    case EntityState.Added:
                        entry.Entity.CreatedDate = DateTime.UtcNow;
                        entry.Entity.CreatedBy = _authenticatedUser == null ? _authenticatedUser.UserId : "system";
                        break;
                    case EntityState.Modified:
                        entry.Entity.LastModifiedDate = DateTime.UtcNow;
                        entry.Entity.LastModifiedBy = _authenticatedUser.UserId;
                        break;
                }
            }
            return base.SaveChangesAsync(cancellationToken);
        }
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<User>(entity =>
            {
                entity.ToTable("user");
                entity.HasKey(_ => _.Id);
                entity.Property(e => e.Id)
                      .HasColumnName("id")
                      .IsRequired();
                entity.Property(e => e.Password)
                      .HasColumnName("password")
                      .IsRequired();
                entity.Property(e => e.Email)
                    .HasColumnName("email")
                    .IsRequired();
                entity.Property(e => e.MaxTodo)
                      .HasColumnName("max_to_do")
                      .IsRequired()
                      .HasDefaultValue(5);
            });
            modelBuilder.Entity<TaskEntity>(entity =>
            {
                entity.ToTable("task");
                entity.HasKey(_ => _.Id);
                entity.Property(e => e.UserId)
                      .HasColumnName("user_id")
                      .IsRequired();
                entity.Property(e => e.Content)
                      .HasColumnName("content")
                      .IsRequired();
                entity.Property(e => e.CreatedDate)
                      .HasColumnName("create_date")
                      .IsRequired();
                entity.HasOne(_ => _.User)
                      .WithMany(_ => _.Tasks)
                      .HasForeignKey(_ => _.UserId)
                      .HasConstraintName("tasks_FK");
            });
        }
    }
}
