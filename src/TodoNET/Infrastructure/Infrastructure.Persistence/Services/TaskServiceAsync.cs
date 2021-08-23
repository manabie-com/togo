using Application.DTOs.Task;
using Application.Exceptions;
using Application.Interfaces;
using Application.Interfaces.Services;
using Application.Wrappers;
using AutoMapper;
using Domain.Entities;
using Infrastructure.Persistence.Contexts;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using TaskEntity = Domain.Entities.Task;

namespace Infrastructure.Persistence.Services
{
    public class TaskServiceAsync : ITaskServiceAsync
    {
        private readonly IGenericRepositoryAsync<TaskEntity> _taskRepositoryAsync;
        private readonly IMapper _mapper;
        private readonly IGenericRepositoryAsync<User> _userRepositoryAsync;
        private readonly ApplicationDbContext _dbContext;
        private readonly IAuthenticatedUserService _authenticatedUser;

        public TaskServiceAsync(IGenericRepositoryAsync<TaskEntity> taskRepositoryAsync, IMapper mapper,
            IGenericRepositoryAsync<User> userRepositoryAsync,
            ApplicationDbContext dbContext,
            IAuthenticatedUserService authenticatedUser)
        {
            _taskRepositoryAsync = taskRepositoryAsync;
            _mapper = mapper;
            _userRepositoryAsync = userRepositoryAsync;
            _dbContext = dbContext;
            _authenticatedUser = authenticatedUser;
        }
        public async Task<Response<TaskResponse>> CreateTaskAsync(CreateTaskRequest request)
        {
            if (request == null)
            {
                throw new ArgumentNullException(nameof(request));
            }
            if (string.IsNullOrEmpty(request.Content))
            {
                throw new ApiException("Invalid input");
            }
            var user = await _userRepositoryAsync.GetByIdAsync(_authenticatedUser.UserId);
            if (user != null)
            {
                var totalTodayTask = _dbContext.Tasks.Where(_ => _.UserId == user.Id && _.CreatedDate.Date == DateTime.UtcNow.Date).Count();
                if(totalTodayTask < user.MaxTodo)
                {
                    var task = _mapper.Map<TaskEntity>(request);
                    task.Id = Guid.NewGuid().ToString();
                    task.UserId = _authenticatedUser.UserId;
                    await _taskRepositoryAsync.AddAsync(task);
                    var result = _mapper.Map<TaskResponse>(task);
                    return new Response<TaskResponse>(result);
                }
                throw new ApiException("Daily limit is reached");
            }
            throw new KeyNotFoundException("Not found this user");
        }

        public  async Task<Response<IReadOnlyList<TaskResponse>>> GetTasksAsync()
        {
            var tasks = await _taskRepositoryAsync.GetAllAsync();
            var result = _mapper.Map<IReadOnlyList<TaskResponse>>(tasks);
            return new Response<IReadOnlyList<TaskResponse>>(result);
        }
    }
}
