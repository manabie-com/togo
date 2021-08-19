using Microsoft.EntityFrameworkCore;
using System;
using System.Linq;
using System.Threading.Tasks;
using WebApi.Models;
using WebApi.Requests;
using WebApi.ViewModels;

namespace WebApi.Services
{
    public class TaskService : ITaskService
    {
        private readonly DemoDbContext _demoDbContext;
        public TaskService(DemoDbContext demoDbContext)
        {
            _demoDbContext = demoDbContext;
        }

        /// <summary>
        ///   Check limit daily task for can add new task or not
        /// </summary>
        /// <param name="userId"></param>
        /// <returns></returns>
        public async Task<bool?> CheckLimitDailyTask(Guid currentUserId)
        {
            var user = await _demoDbContext.Users.SingleOrDefaultAsync(x => x.Id == currentUserId);
            if (user == null) return null;

            var tasksForToday = _demoDbContext.Tasks.Count(x => x.UserId == currentUserId && x.CreatedDate.Date == DateTime.Now.Date);

            return user.MaxTodo > tasksForToday;
        }

        /// <summary>
        /// Create new task for current user request
        /// </summary>
        /// <param name="createTaskRequest"></param>
        /// <param name="currentUserId"></param>
        /// <returns></returns>
        public async Task<CreateTaskViewModel> CreateTask(CreateTaskRequest createTaskRequest, Guid currentUserId)
        {
            var newTask = new Models.Task
            {
                Id = Guid.NewGuid(),
                UserId = currentUserId,
                Content = createTaskRequest.Content.Trim(),
                CreatedDate = DateTime.Now
            };
            await _demoDbContext.AddAsync(newTask);
            await _demoDbContext.SaveChangesAsync();

            return new CreateTaskViewModel
            {
                Id = newTask.Id,
                Content = newTask.Content
            };
        }
    }
}
