using System.Reflection;
using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;
using Togo.Core.Entities;
using Togo.Core.Interfaces;
using Togo.Infrastructure.Identities;

namespace Togo.Infrastructure.Persistence;

// dotnet ef migrations add InitialCreate -p Togo.Infrastructure -s Togo.Api -o Persistence/Scripts

public class AppDbContext : IdentityDbContext<AppUser>
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
}
