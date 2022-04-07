using System;

using Manabie.Api.Insfrastructures;
using Manabie.Api.Models;

using Microsoft.AspNetCore.Identity;

namespace Manabie.Api.Services
{
    public class TaskService : ITaskService
    {
        private readonly ManaDbContext _context;
        public TaskService(ManaDbContext manaDbContext)
        {
            _context = manaDbContext;
        }

        public IEnumerable<TaskListViewModel> GetTasks(string userName)
        {
            var user = GetUserByUserName(userName);
            if (user == null) return null;
            return _context.Tasks.Where(n => n.UserId == user.Id)
                .Select(n => new TaskListViewModel
                {
                    Todo = n.Todo
                }).ToList();
        }

        public MessageResponse AddTask(string userName, TaskViewModel task)
        {
            if (string.IsNullOrEmpty(task.Todo))
            {
                return new MessageResponse
                {
                    Success = false,
                    Message = "Todo can't be empty."
                };
            }

            var user = GetUserByUserName(userName);
            if (user == null)
            {
                return new MessageResponse
                {
                    Success = false,
                    Message = "This user is not exist in the system."
                };
            }
            var countTask = _context.Tasks.Count(n => n.UserId == user.Id && n.CreatedAt.Date == DateTime.UtcNow.Date);
            if (countTask == user.MaxTodo)
            {
                return new MessageResponse
                {
                    Success = false,
                    Message = "Your list is maximized."
                };
            }

            _context.Tasks.Add(new Entities.Task { Todo = task.Todo, UserId = user.Id });
            return new MessageResponse
            {
                Success = _context.SaveChanges() > 0
            };
        }

        private Entities.User GetUserByUserName(string userName)
        {
            return _context.Users.SingleOrDefault(n => n.Username == userName);
        }
    }
}

