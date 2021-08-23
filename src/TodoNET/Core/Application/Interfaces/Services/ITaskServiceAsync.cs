using Application.DTOs.Task;
using Application.Wrappers;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Application.Interfaces.Services
{
    public interface ITaskServiceAsync
    {
        Task<Response<TaskResponse>> CreateTaskAsync(CreateTaskRequest request);
        Task<Response<IReadOnlyList<TaskResponse>>> GetTasksAsync();
    }
}
