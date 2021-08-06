using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using togo.Service.Context;
using togo.Service.Dto;
using togo.Service.Interface;
using TaskEntity = togo.Service.Context.Task;
using System.Linq;
using Microsoft.EntityFrameworkCore;
using togo.Service.Errors;
using System.Net;

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
            if (string.IsNullOrEmpty(input.Content))
            {
                throw new RestException(HttpStatusCode.BadRequest);
            }

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

        public async Task<List<TaskEntity>> List(string created_date)
        {
            var query = from t in _context.Tasks
                        where t.UserId == _currentHttpContext.GetCurrentUserId()
                        select t;

            if (!string.IsNullOrEmpty(created_date))
            {
                bool canParse = DateTime.TryParse(created_date, out var date);
                if (!canParse)
                {
                    throw new RestException(HttpStatusCode.BadRequest);
                }

                query = from t in query
                        where t.CreatedDate == date.ToShortDateString()
                        select t;
            }

            return await query.ToListAsync();
        }
    }
}
