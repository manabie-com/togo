using Microsoft.EntityFrameworkCore;
using Togo.Core.Entities;

namespace Togo.Core.Interfaces;

public interface IAppDbContext
{
    DbSet<TaskItem> TaskItems { get; set; }

    int SaveChanges();

    Task<int> SaveChangesAsync(CancellationToken cancellationToken = new());
}
