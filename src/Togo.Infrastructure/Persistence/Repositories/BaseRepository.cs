using System.Linq.Expressions;
using Togo.Core.Entities.Common;
using Togo.Core.Interfaces;
using Togo.Core.Interfaces.Repositories;

namespace Togo.Infrastructure.Persistence.Repositories;

public abstract class BaseRepository<TEntity> : IRepository<TEntity> where TEntity : BaseEntity
{
    protected readonly ICurrentUserService CurrentUserService;

    protected BaseRepository(ICurrentUserService currentUserService)
    {
        CurrentUserService = currentUserService;
    }

    public async Task<TEntity> AddAsync(TEntity entity)
    {
        SetCreationAuditedFields(entity);
        return await CreateEntityAsync(entity);
    }

    public abstract Task<int> CountAsync(Expression<Func<TEntity, bool>> expression);

    protected abstract Task<TEntity> CreateEntityAsync(TEntity entity);

    protected void SetCreationAuditedFields(TEntity entity)
    {
        if (typeof(DateEntity).IsAssignableFrom(typeof(TEntity)))
        {
            if (entity is DateEntity dateEntity)
            {
                dateEntity.CreatedAt = DateTime.UtcNow;
            }
        }

        if (typeof(AuditedEntity).IsAssignableFrom(typeof(TEntity)))
        {
            if (entity is AuditedEntity auditedEntity)
            {
                auditedEntity.CreatedBy = CurrentUserService.UserId;    
            }
        }
    }
}
