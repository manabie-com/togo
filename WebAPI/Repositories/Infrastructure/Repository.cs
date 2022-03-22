using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Text;
using System.Threading.Tasks;

namespace Repositories.Infrastructure
{
    public class Repository<T> : IRepository<T>, IDisposable where T : class
    {
        private DatabaseContext _context;

        public Repository(DatabaseContext context)
        {
            _context = context;
            _dbSet = _context.Set<T>();
        }

        private DbSet<T> _dbSet;

        public DbSet<T> DbSet
        {
            get
            {
                if (_dbSet == null)
                {
                    _dbSet = _context.Set<T>();
                }
                return _dbSet as DbSet<T>;
            }
        }

        /// <summary>
        /// find single
        /// </summary>
        /// <param name="predicate"></param>
        /// <param name="includeProperties"></param>
        /// <returns></returns>
        public virtual T FindSingle(Expression<Func<T, bool>> predicate)
        {
            return FindAll().FirstOrDefault(predicate);
        }

        /// <summary>
        /// find all
        /// </summary>
        /// <param name="includeProperties"></param>
        /// <returns></returns>
        public virtual IQueryable<T> FindAll()
        {
            IQueryable<T> items = _context.Set<T>();
            
            return items;
        }

        /// <summary>
        /// find all with condition
        /// </summary>
        /// <param name="predicate"></param>
        /// <param name="includeProperties"></param>
        /// <returns></returns>
        public virtual IQueryable<T> FindAll(Expression<Func<T, bool>> predicate)
        {
            IQueryable<T> items = _context.Set<T>();
            
            return items.Where(predicate);
        }

        /// <summary>
        /// count
        /// </summary>
        /// <param name="predicate"></param>
        /// <returns></returns>
        public virtual int Count(Expression<Func<T, bool>> predicate)
        {
            return _dbSet.Count(predicate);
        }

        /// <summary>
        /// add
        /// </summary>
        /// <param name="entity"></param>
        public virtual void Add(T entity)
        {
            _dbSet.Add(entity);
        }

        /// <summary>
        /// add multi
        /// </summary>
        /// <param name="listEntity"></param>
        public virtual void AddMulti(IList<T> listEntity)
        {
            _dbSet.AddRange(listEntity);
        }

        /// <summary>
        /// update
        /// </summary>
        /// <param name="entity"></param>
        public virtual void Update(T entity)
        {
            _dbSet.Update(entity);
        }

        /// <summary>
        /// update multi
        /// </summary>
        /// <param name="listEntity"></param>
        public virtual void UpdateMulti(IList<T> listEntity)
        {
            _dbSet.UpdateRange(listEntity);
        }

        /// <summary>
        /// remove
        /// </summary>
        /// <param name="entity"></param>
        public virtual void Remove(T entity)
        {
            _dbSet.Remove(entity);
        }

        /// <summary>
        /// remove multi
        /// </summary>
        /// <param name="listEntity"></param>
        public virtual void RemoveMulti(IList<T> listEntity)
        {
            _dbSet.RemoveRange(listEntity);
        }

        /// <summary>
        /// begin transaction
        /// </summary>
        public void BeginTransaction()
        {
            _context.Database.BeginTransaction();
        }

        /// <summary>
        /// commit transaction
        /// </summary>
        public void CommitTransaction()
        {
            _context.Database.CommitTransaction();
        }

        /// <summary>
        /// rollback transaction
        /// </summary>
        public void RollbackTransaction()
        {
            _context.Database.RollbackTransaction();
        }

        public void Dispose()
        {
            if (_context != null)
            {
                _context.Dispose();
            }
        }
    }
}
