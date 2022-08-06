using System.Linq.Expressions;
using Microsoft.EntityFrameworkCore;
using Togo.Core.Entities.Common;
using Togo.Core.Interfaces;

namespace Togo.Infrastructure.Persistence.Repositories;

public class EfRepository<TEntity> : BaseRepository<TEntity> where TEntity : BaseEntity
{
    protected readonly AppDbContext DbContext;
        
    public EfRepository(
        ICurrentUserService currentUserService, 
        AppDbContext dbContext) : base(currentUserService)
    {
        DbContext = dbContext;
    }

    public override async Task<int> CountAsync(Expression<Func<TEntity, bool>> expression)
    {
        return await DbContext.Set<TEntity>()
            .CountAsync(expression);
    }

    protected override async Task<TEntity> CreateEntityAsync(TEntity entity)
    {
        var added = await DbContext
            .Set<TEntity>()
            .AddAsync(entity);
            
        return added.Entity;
    }
}
