using System;
using System.Threading.Tasks;
using WebApi.Requests;
using WebApi.ViewModels;

namespace WebApi.Services
{
    public interface ITaskService
    {
        /// <summary>
        /// Create new task for current user request
        /// </summary>
        /// <param name="createTaskRequest"></param>
        /// <param name="currentUserId"></param>
        /// <returns></returns>
        public Task<CreateTaskViewModel> CreateTask(CreateTaskRequest createTaskRequest, Guid currentUserId);

        /// <summary>
        ///   Check limit daily task for can add new task or not
        /// </summary>
        /// <param name="currentUserId"></param>
        /// <returns></returns>
        public Task<bool?> CheckLimitDailyTask(Guid currentUserId);
    }
}
