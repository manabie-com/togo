using System.Threading.Tasks;
using TogoService.API.Infrastructure.Database;
using TogoService.API.Infrastructure.Repository;

namespace TogoService.API.Model.Interface
{
    public interface IUnitOfWork
    {
        BaseRepository<T> GenericRepository<T>() where T : class;
        TogoDbContext Context { get; }
        Task CreateTransactionAsync();
        Task CommitAsync();
        Task RollbackAsync();
        Task<int> Save();
    }
}