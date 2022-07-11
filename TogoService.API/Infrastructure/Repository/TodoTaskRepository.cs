using System;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using TogoService.API.Model;
using TogoService.API.Model.Interface;

namespace TogoService.API.Infrastructure.Repository
{
    public class TodoTaskRepository : BaseRepository<TodoTask>, ITodoTaskRepository
    {
        public TodoTaskRepository() { }

        public async Task<TodoTask[]> GetAddedTasks(Guid userId, DateTime todoDate)
        {
            try
            {
                return await (_dbContext.TodoTasks.Where(x => x.IsDeleted == false
                    && x.UserId.Equals(userId)
                    && x.TodoDay.Date.Equals(todoDate.Date)).ToArrayAsync());
            }
            catch (Exception ex)
            {
                throw ex;
            }
        }

        public void SetUnitOfWork(IUnitOfWork unitOfWork)
        {
            this._dbContext = unitOfWork.Context;
        }
    }
}
