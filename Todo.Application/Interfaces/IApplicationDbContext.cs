using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Todo.Domain.Entities;

namespace Todo.Application.Interfaces
{
    public interface IApplicationDbContext : IDisposable
    {
        DatabaseFacade Database { get; }
        DbSet<User> Users { get; set; }
        DbSet<UserTask> UserTasks { get; set; }
        int SaveChanges();
        Task<int> SaveChangesAsync(CancellationToken cancellationToken = default);

        Task CommitAsync(Func<Task> action);
    }
}

