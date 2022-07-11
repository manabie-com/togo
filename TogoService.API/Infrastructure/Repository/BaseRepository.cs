using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using TogoService.API.Infrastructure.Database;
using TogoService.API.Model.Interface;

namespace TogoService.API.Infrastructure.Repository
{
    public class BaseRepository<T> : IRepository<T>, IDisposable where T : class
    {
        private DbSet<T> _entities;
        protected TogoDbContext _dbContext;

        public BaseRepository() { }

        public BaseRepository(TogoDbContext dbContext)
        {
            _dbContext = dbContext;
        }

        public BaseRepository(IUnitOfWork unitOfWork)
            : this(unitOfWork.Context) { }

        protected virtual DbSet<T> Entities
        {
            get { return _entities ?? (_entities = _dbContext.Set<T>()); }
        }

        public void Dispose()
        {
            if (_dbContext != null)
                _dbContext.Dispose();
        }

        public virtual async Task<T> GetById(Guid id)
        {
            try
            {
                return await Entities.FindAsync(id);
            }
            catch (Exception ex)
            {
                throw ex;
            }
        }

        public virtual async Task<T> Add(T entity)
        {
            try
            {
                if (entity == null)
                    throw new ArgumentNullException(typeof(T).Name);
                return (await Entities.AddAsync(entity)).Entity;
            }
            catch (Exception ex)
            {
                throw ex;
            }
        }

        public virtual async Task AddRange(T[] entities)
        {
            try
            {
                await _dbContext.AddRangeAsync(entities);
            }
            catch (Exception ex)
            {
                throw ex;
            }
        }

        public void HardDelete(T entity)
        {
            throw new NotImplementedException();
        }

        public void HardDeleteRange(T[] entities)
        {
            throw new NotImplementedException();
        }

        public void Update(T entity)
        {
            throw new NotImplementedException();
        }

        public void UpdateRange(T[] entities)
        {
            throw new NotImplementedException();
        }

        public Task<IEnumerable<T>> GetAll()
        {
            throw new NotImplementedException();
        }
    }
}
