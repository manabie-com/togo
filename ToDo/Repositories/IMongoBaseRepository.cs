using MongoDB.Driver;
using ToDo.Api.Domain.Core;

namespace ToDo.Api.Repositories
{
    public interface IMongoBaseRepository<T> where T : IEntity
    {
        Task AddAsync(T obj, CancellationToken cancellationToken);
        Task<List<T>> FindAsync(CancellationToken cancellationToken);
        Task<List<T>> FindByIdAsync(Guid id, CancellationToken cancellationToken);
        Task UpdateAsync(T obj, CancellationToken cancellationToken);
        Task<T> GetByIdAsync(Guid id, CancellationToken cancellationToken);
        Task DeleteAsync(Guid id, CancellationToken cancellationToken);
        Task<IClientSessionHandle> StartSessionAsync(CancellationToken cancellationToken);
        Task<List<T>> FindAllWithFilter(FilterDefinition<T> filter);
    }
}