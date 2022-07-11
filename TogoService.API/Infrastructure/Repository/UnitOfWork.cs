using TogoService.API.Model.Interface;
using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using Microsoft.EntityFrameworkCore.Storage;
using TogoService.API.Infrastructure.Database;

namespace TogoService.API.Infrastructure.Repository
{
    public class UnitOfWork : IUnitOfWork, IDisposable
    {
        private readonly TogoDbContext _dbContext;
        private IDbContextTransaction _objTran;
        private bool _disposed;
        private Dictionary<string, object> _repositories;

        public UnitOfWork(TogoDbContext dbContext)
        {
            _dbContext = dbContext;
        }

        public TogoDbContext Context
        {
            get { return _dbContext; }
        }

        public void Dispose()
        {
            Dispose(true);
            GC.SuppressFinalize(this);
        }

        public BaseRepository<T> GenericRepository<T>() where T : class
        {
            if (_repositories == null)
            {
                _repositories = new Dictionary<string, object>();
            }

            var type = typeof(T).Name;

            if (!_repositories.ContainsKey(type))
            {
                var repositoryInstance = new BaseRepository<T>(_dbContext);
                _repositories.Add(type, repositoryInstance);
            }

            return (BaseRepository<T>)_repositories[type];
        }

        protected virtual void Dispose(bool disposing)
        {
            if (!_disposed)
                if (disposing)
                    _dbContext.Dispose();
            _disposed = true;
        }


        public async Task CreateTransactionAsync()
        {
            _objTran = await _dbContext.Database.BeginTransactionAsync();
        }

        public async Task CommitAsync()
        {
            await _objTran.CommitAsync();
        }

        public async Task RollbackAsync()
        {
            await _objTran.RollbackAsync();
            await _objTran.DisposeAsync();
        }

        public async Task<int> Save()
        {
            try
            {
                return await _dbContext.SaveChangesAsync();
            }
            catch (Exception ex)
            {
                throw ex;
            }
        }
    }
}
