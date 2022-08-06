using Togo.Core.Entities;
using Togo.Core.Entities.Common;
using Togo.Core.Interfaces.Repositories;

namespace Togo.Core.Interfaces;

public interface IUnitOfWork
{
    IRepository<TEntity> GetRepository<TEntity>() where TEntity : BaseEntity;

    IRepository<TaskItem> TaskItemRepository { get; }

    Task CommitAsync();
}
