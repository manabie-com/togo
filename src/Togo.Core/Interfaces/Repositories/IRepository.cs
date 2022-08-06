using System.Linq.Expressions;
using Togo.Core.Entities.Common;

namespace Togo.Core.Interfaces.Repositories;

public interface IRepository<TEntity> where TEntity : BaseEntity
{
    Task<TEntity> AddAsync(TEntity entity);

    Task<int> CountAsync(Expression<Func<TEntity, bool>> expression);
}
