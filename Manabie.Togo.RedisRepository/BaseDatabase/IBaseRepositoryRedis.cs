using Manabie.Togo.Data.Base;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Manabie.Togo.RedisRepository.BaseDatabase
{
    public interface IBaseRepositoryRedis<T> where T : BaseEntity
    {
        T GetByID(Guid id);
        Task<T> GetByIDAsync(Guid id);
        IEnumerable<T> GetMultiple(IEnumerable<Guid> ids);
        Task<IEnumerable<T>> GetMultipleAsync(IEnumerable<Guid> ids);
        IEnumerable<T> GetAll();
        Task<IEnumerable<T>> GetAllAsync();
        bool Exists(Guid id);
        bool Exists(string key);
        void Save(T item);
        Task SaveAsync(T item);
        void Save(IEnumerable<T> items);
        Task SaveAsync(IEnumerable<T> items);
        void Remove(T item);
        Task RemoveAsync(T item);
        void Remove(IEnumerable<T> items);
        Task RemoveAsync(IEnumerable<T> items);
        void Remove(Guid id);
        Task RemoveAsync(Guid id);
        void Remove(IEnumerable<Guid> ids);
        Task RemoveAsync(IEnumerable<Guid> ids);
        void Remove(string key);
        Task RemoveAsync(string key);
        void Remove(IEnumerable<string> keys);
        Task RemoveAsync(IEnumerable<string> keys);

        void ClearAll();
        Task ClearAllAsync();
    }
}
