using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace TogoService.API.Model.Interface
{
    public interface IRepository<T> where T : class
    {
        /// <summary>
        /// Add an entity
        /// </summary>
        /// <param name="entity">Entity</param>
        /// <returns>Entity after added</returns>
        Task<T> Add(T entity);

        /// <summary>
        /// Add a list of entity
        /// </summary>
        /// <param name="entities">List of entity</param>
        Task AddRange(T[] entities);

        /// <summary>
        /// Hard delete an entity
        /// </summary>
        /// <param name="entity">Entity</param>
        void HardDelete(T entity);

        /// <summary>
        /// Hard delete a list of entity
        /// </summary>
        /// <param name="entities">List of entity</param>
        void HardDeleteRange(T[] entities);

        /// <summary>
        /// Change state of entity to modified
        /// </summary>
        /// <param name="entity">Entity</param>
        void Update(T entity);

        /// <summary>
        /// Change state of entities to modified
        /// </summary>
        /// <param name="entities">entities</param>
        void UpdateRange(T[] entities);

        /// <summary>
        /// Get entity by its id.
        /// </summary>
        /// <param name="id">Guid</param>
        /// <returns>Entity</returns>
        Task<T> GetById(Guid id);

        /// <summary>
        /// Get all entities.
        /// </summary>
        /// <returns>List of entitties</returns>
        Task<IEnumerable<T>> GetAll();
    }
}
