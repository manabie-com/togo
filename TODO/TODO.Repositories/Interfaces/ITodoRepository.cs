using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using TODO.Repositories.Data.DBModels;
using TODO.Repositories.Models.RequestModels;

namespace TODO.Repositories.Interfaces
{
    public interface ITodoRepository
    {
        Task<Todo> CreateTodo(CreateTodoRequest request);
    }
}
