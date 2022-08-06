using Togo.Core.Entities;
using Togo.Core.Entities.Common;
using Togo.Core.Interfaces;
using Togo.Core.Interfaces.Repositories;
using Togo.Infrastructure.Persistence.Repositories;

namespace Togo.Infrastructure.Persistence;

public class EfUnitOfWork : IUnitOfWork
{
    private readonly AppDbContext _dbContext;
    private readonly ICurrentUserService _currentUserService;

    public EfUnitOfWork(AppDbContext dbContext, ICurrentUserService currentUserService)
    {
        _dbContext = dbContext;
        _currentUserService = currentUserService;
    }
    
    private IRepository<TaskItem>? _taskItemRepository;
    public IRepository<TaskItem> TaskItemRepository =>
        _taskItemRepository ??= new EfRepository<TaskItem>(_currentUserService, _dbContext);

    public IRepository<TEntity> GetRepository<TEntity>() where TEntity : BaseEntity
    {
        if (typeof(TEntity) == typeof(TaskItem))
        {
            return (IRepository<TEntity>) TaskItemRepository;
        }

        throw new InvalidOperationException(
            $"Unable to get repository for entity type {typeof(TEntity)}");
    }

    public async Task CommitAsync()
    {
        await _dbContext.SaveChangesAsync();
    }
}
