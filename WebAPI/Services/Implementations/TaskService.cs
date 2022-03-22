using Microsoft.Extensions.Logging;
using Models;
using Repositories.Infrastructure;
using Services.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Services.Implementations
{
    public class TaskService : ITaskService
    {
        private readonly IRepository<Tasks> _taskRepository;
        private readonly IUnitOfWork _unitOfWork;
        private readonly IRepository<Users> _userRepository;
        private readonly ILogger _logger;

        public TaskService(IRepository<Tasks> taskRepository,
                           IUnitOfWork unitOfWork,
                           IRepository<Users> userRepository,
                           ILoggerFactory loggerFactory)
        {
            _taskRepository = taskRepository;
            _unitOfWork = unitOfWork;
            _logger = loggerFactory.CreateLogger("TaskService");
            _userRepository = userRepository;
        }

        public int Create(Tasks task, string userId)
        {
            try
            {
                var taskPerDay = _userRepository.FindSingle(_ => _.Id == userId).TaskPerDay;

                var min = new DateTime(DateTime.Now.Year, DateTime.Now.Month, DateTime.Now.Day, 0, 0, 0);
                var max = new DateTime(DateTime.Now.Year, DateTime.Now.Month, DateTime.Now.Day, 23, 59, 59);
                var countExistTask = _taskRepository.Count(_ => _.UserID == userId && _.CreateAt >= min && _.CreateAt <= max);

                if(countExistTask < taskPerDay)
                {
                    task.ID = Guid.NewGuid().ToString();
                    task.CreateAt = DateTime.Now;
                    task.UserID = userId;
                    _taskRepository.Add(task);
                    _unitOfWork.Commit();
                    return 1;
                }

                return 0;
            }
            catch(Exception ex)
            {
                _logger.LogError("TaskService.Create: " + ex.ToString());
                return -1;
            }
        }

        public List<Tasks> GetTasksByUserId(string userId)
        {
            return _taskRepository.FindAll(_ => _.UserID == userId).ToList();
        }
    }
}
