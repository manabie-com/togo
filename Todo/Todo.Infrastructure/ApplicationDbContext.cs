using Microsoft.EntityFrameworkCore;
using System.Diagnostics;
using Todo.Application.Interfaces;
using Todo.Domain;
using Todo.Domain.Entities;
using Todo.Infrastructure.EntityConfigurations;

namespace Todo.Infrastructure
{
    public class ApplicationDbContext : DbContext, IApplicationDbContext
    {
        public DbSet<User> Users { get; set; }
        public DbSet<UserTask> UserTasks { get; set; }

        public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options) : base(options)
        {
        }
        protected override void OnModelCreating(ModelBuilder builder)
        {
            builder.HasDefaultSchema("public");

            base.OnModelCreating(builder);

            builder.ApplyConfigurationsFromAssembly(typeof(UserConfiguration).Assembly);
        }
        public override int SaveChanges()
        {
            BeforeCommit();
            return base.SaveChanges();
        }
        public override Task<int> SaveChangesAsync(CancellationToken cancellationToken = default)
        {
            BeforeCommit();
            return base.SaveChangesAsync(cancellationToken);
        }
        private void BeforeCommit()
        {
            var entriesAdded = ChangeTracker.Entries()
                .Where(e => e.State == EntityState.Added)
                .Select(e => e.Entity);

            var entriesModified = ChangeTracker.Entries()
                  .Where(e => e.State == EntityState.Modified).Select(e => e.Entity as IAuditableEntity);

            if (entriesAdded.Any()) ProcessAudit(entriesAdded, EntityState.Added);

            if (entriesModified.Any()) ProcessAudit(entriesModified, EntityState.Modified);
        }

        private void ProcessAudit(IEnumerable<object> entries, EntityState state)
        {
            foreach (var e in entries.Select(e => e as IAuditableEntity))
            {
                if (e is not null)
                {
                    if (state == EntityState.Added)
                    {
                        //e.CreatedBy = _userId;
                        e.CreatedAt = DateTime.UtcNow;
                    }
                    else
                    {
                        e.ModifiedAt = DateTime.UtcNow;
                    }
                }
            }
        }
        public async Task CommitAsync(Func<Task> action)
        {
            var strategy = Database.CreateExecutionStrategy();

            await strategy.ExecuteAsync(async () =>
            {
                using var transaction = Database.BeginTransaction();

                try
                {
                    await base.SaveChangesAsync();

                    await action();
                }
                catch (Exception ex)
                {
                    Trace.TraceError(ex.Message, ex);

                    transaction.Rollback();

                    throw;
                }
                finally
                {
                    transaction.Commit();
                }
            });
        }
    }
}
