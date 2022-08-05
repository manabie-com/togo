using System.Reflection;
using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;
using Togo.Core.Entities;
using Togo.Core.Entities.Common;
using Togo.Core.Interfaces;
using Togo.Infrastructure.Identities;

namespace Togo.Infrastructure.Persistence;

// dotnet ef migrations add InitialCreate -p Togo.Infrastructure -s Togo.Api -o Persistence/Scripts
// dotnet ef migrations add CreateTaskItemEntity -p Togo.Infrastructure -s Togo.Api -o Persistence/Scripts

public class AppDbContext : IdentityDbContext<AppUser>, IAppDbContext
{
    private readonly ICurrentUserService _currentUserService;
    
    public DbSet<TaskItem> TaskItems { get; set; }

    public AppDbContext(DbContextOptions<AppDbContext> options, ICurrentUserService currentUserService) : base(options)
    {
        _currentUserService = currentUserService;
    }

    protected override void OnModelCreating(ModelBuilder builder)
    {
        base.OnModelCreating(builder);
        builder.ApplyConfigurationsFromAssembly(Assembly.GetExecutingAssembly());
    }

    public override int SaveChanges()
    {
        foreach (var entry in ChangeTracker.Entries<BaseEntity>())  
        {
            if (entry.State is EntityState.Added or EntityState.Modified)  
            {
                SetModificationAuditedFields(entry.Entity);

                if (entry.State == EntityState.Added)
                {
                    SetCreationAuditedFields(entry.Entity);
                }
            }
        }
        
        return base.SaveChanges();
    }

    public override Task<int> SaveChangesAsync(CancellationToken cancellationToken = new CancellationToken())
    {
        foreach (var entry in ChangeTracker.Entries<DateEntity>())  
        {
            if (entry.State is EntityState.Added or EntityState.Modified)  
            {
                SetModificationAuditedFields(entry.Entity);
                
                if (entry.State == EntityState.Added)
                {
                    SetCreationAuditedFields(entry.Entity);
                }
            }
        }
        return base.SaveChangesAsync(cancellationToken);
    }
    
    private void SetCreationAuditedFields(BaseEntity entity)
    {
        if (entity is DateEntity dateEntity)
        {
            dateEntity.CreatedAt = DateTime.UtcNow;
        }

        if (entity is AuditedEntity auditedEntity)
        {
            auditedEntity.CreatedBy = _currentUserService.UserId;    
        }
    }
    
    private void SetModificationAuditedFields(BaseEntity entity)
    {
        if (entity is DateEntity dateEntity)
        {
            dateEntity.LastModifiedAt = DateTime.UtcNow;
        }

        if (entity is AuditedEntity auditedEntity)
        {
            auditedEntity.LastModifiedBy = _currentUserService.UserId;    
        }
    }
}
