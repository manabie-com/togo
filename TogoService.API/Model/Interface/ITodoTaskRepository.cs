using System;
using System.Threading.Tasks;

namespace TogoService.API.Model.Interface
{
    public interface ITodoTaskRepository : IRepository<TodoTask>
    {
        void SetUnitOfWork(IUnitOfWork unitOfWork);

        /// <summary>
        /// Get tasks already added for one user on specific date.
        /// </summary>
        /// <param name="userId"></param>
        /// <param name="todoDate"></param>
        /// <returns>List tasks already added.</returns>
        Task<TodoTask[]> GetAddedTasks(Guid userId, DateTime todoDate);
    }
}