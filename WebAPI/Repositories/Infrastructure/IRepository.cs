using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Text;
using System.Threading.Tasks;

namespace Repositories.Infrastructure
{
    public interface IRepository<T> where T : class
    {
        /// <summary>
        /// DbSet property
        /// </summary>
        DbSet<T> DbSet { get; }

        /// <summary>
        /// find single
        /// </summary>
        /// <param name="predicate"></param>
        /// <param name="includeProperties"></param>
        /// <returns></returns>
        T FindSingle(Expression<Func<T, bool>> predicate);

        /// <summary>
        /// find all
        /// </summary>
        /// <param name="includeProperties"></param>
        /// <returns></returns>
        IQueryable<T> FindAll();

        /// <summary>
        /// find all with condition
        /// </summary>
        /// <param name="predicate"></param>
        /// <param name="includeProperties"></param>
        /// <returns></returns>
        IQueryable<T> FindAll(Expression<Func<T, bool>> predicate);

        /// <summary>
        /// count
        /// </summary>
        /// <param name="predicate"></param>
        /// <returns></returns>
        int Count(Expression<Func<T, bool>> predicate);

        /// <summary>
        /// add
        /// </summary>
        /// <param name="entity"></param>
        void Add(T entity);

        /// <summary>
        /// add multi
        /// </summary>
        /// <param name="listEntity"></param>
        void AddMulti(IList<T> listEntity);

        /// <summary>
        /// update
        /// </summary>
        /// <param name="entity"></param>
        void Update(T entity);

        /// <summary>
        /// update multi
        /// </summary>
        /// <param name="listEntity"></param>
        void UpdateMulti(IList<T> listEntity);

        /// <summary>
        /// remove
        /// </summary>
        /// <param name="entity"></param>
        void Remove(T entity);

        /// <summary>
        /// remove multi
        /// </summary>
        /// <param name="listEntity"></param>
        void RemoveMulti(IList<T> listEntity);

        /// <summary>
        /// begin transaction
        /// </summary>
        void BeginTransaction();

        /// <summary>
        /// rollback transaction
        /// </summary>
        void RollbackTransaction();

        /// <summary>
        /// commit transaction
        /// </summary>
        void CommitTransaction();
    }
}
