using System;
using System.Threading.Tasks;
using togo.Service.Context;
using togo.Service.Dto;
using togo.Service.Interface;
using TaskEntity = togo.Service.Context.Task;

namespace togo.Service
{
    public class TaskService : ITaskService
    {
        private readonly TogoContext _context;
        private readonly ICurrentHttpContext _currentHttpContext;

        public TaskService(
              TogoContext context
            , ICurrentHttpContext currentHttpContext)
        {
            _context = context;
            _currentHttpContext = currentHttpContext;
        }

        public async Task<TaskEntity> Create(TaskCreateDto input)
        {
            var task = new TaskEntity
            {
                Id = Guid.NewGuid().ToString(),
                Content = input.Content,
                UserId = _currentHttpContext.GetCurrentUserId(),
                CreatedDate = DateTime.Now.ToShortDateString(),
            };

            await _context.AddAsync(task);
            await _context.SaveChangesAsync();

            return task;
        }
    }
}
