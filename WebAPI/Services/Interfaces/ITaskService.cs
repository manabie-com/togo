using Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Services.Interfaces
{
    public interface ITaskService
    {
        int Create(Tasks task, string userId);
        List<Tasks> GetTasksByUserId(string userId);
    }
}
