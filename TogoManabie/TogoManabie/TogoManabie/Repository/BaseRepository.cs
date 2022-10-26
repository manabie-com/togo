using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using TogoManabie.Configuration;
using TogoManabie.Interfaces;
using TogoManabie.Models;

namespace TogoManabie.Repository
{
    public class BaseRepository<T> : IRepository<T> where T : BaseModel
    {
        private readonly EntityFrameworkSqlServerContext _dbContext;

        public BaseRepository(EntityFrameworkSqlServerContext dbContext)
        {
            this._dbContext = dbContext;
        }
        public virtual void Create(T entity)
        {
            _dbContext.Set<T>().Add(entity);
        }

        public async Task<T> GetById(int id)
        {
            return await _dbContext.Set<T>()
                .AsNoTracking()
                .FirstOrDefaultAsync(s => s.id == id);
        }

        
    }
}
