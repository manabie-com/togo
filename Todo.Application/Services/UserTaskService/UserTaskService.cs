using AutoMapper;
using Todo.Application.Dtos;
using Todo.Application.Interfaces;
using Todo.Domain.Entities;

namespace Todo.Application.Services
{
    public interface IUserTaskService
    {
        Task<string> CreateTaskAsync(CreateEditTaskDto dto);
    }
    public class UserTaskService
    : BaseService, IUserTaskService
    {
        public UserTaskService(IMapper mapper, IApplicationDbContext dbContext) : base(mapper, dbContext) { }
        public async Task<string> CreateTaskAsync(CreateEditTaskDto dto)
        {
            var task = Mapper.Map<UserTask>(dto);

            var dateRange = GetCurrentDateRange(DateTime.UtcNow);
            var user = GetUserById(task.CreatedBy);
            var cntTasks = TotalCurrentTask(dateRange.StartDate, dateRange.EndDate, task.CreatedBy);

            if (IsSmallerLimitTask(user.LimitTask, cntTasks))
            {
                await CreateNewTask(task);
            }
            return task.Id ?? String.Empty;
        }

        public bool IsSmallerLimitTask(int limit, int count)
        {
            if (limit >= count + 1)
            {
                return true;
            }
            return false;

        }
        public User GetUserById(string userId)
        {
            var user = DbContext.Users.Where(w => w.Id == userId).SingleOrDefault();
            if (user is null)
            {
                throw new ArgumentException("User is not exists");
            }
            return user;
        }
        public int TotalCurrentTask(DateTime startDate, DateTime endDate, string userId)
        {
            int count = DbContext.UserTasks.Where(w => w.CreatedBy == userId
    && w.CreatedAt >= startDate && w.CreatedAt <= endDate).Count();

            return count;
        }
        public (DateTime StartDate, DateTime EndDate) GetCurrentDateRange(DateTime date)
        {
            var startDate = date.Date;///mm/yyyy 00:00:00
            var endDate = startDate.AddDays(1).AddTicks(-1L);// dd/MM/yyyy 23:59:59

            return (startDate, endDate);
        }
        public async Task CreateNewTask(UserTask task)
        {
            DbContext.UserTasks.Add(task);

            DbContext.SaveChanges();

            await Task.CompletedTask;

        }
    }
}
