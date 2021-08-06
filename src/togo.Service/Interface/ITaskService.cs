using System.Collections.Generic;
using System.Threading.Tasks;
using togo.Service.Dto;
using TaskEntity = togo.Service.Context.Task;

namespace togo.Service.Interface
{
    public interface ITaskService
    {
        Task<TaskEntity> Create(TaskCreateDto input);
        Task<List<TaskEntity>> List(string created_date);
    }
}
