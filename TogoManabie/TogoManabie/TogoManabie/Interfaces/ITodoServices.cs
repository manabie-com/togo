using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using TogoManabie.Models;

namespace TogoManabie.Interfaces
{
    public interface ITodoServices
    {
        Task<Tasks> CreateTodo(Tasks task);
    }
}
