using System;

using Manabie.Api.Models;

namespace Manabie.Api.Services
{
    public interface ITaskService
    {
        IEnumerable<TaskListViewModel> GetTasks(string userName);
        MessageResponse AddTask(string userName, TaskViewModel task);
    }
}

