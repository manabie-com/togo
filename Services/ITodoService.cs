using ManabieTodo.Models;

namespace ManabieTodo.Services
{
    public interface ITodoService
    {
        TodoModel Get(int id);
        Task<TodoModel> GetAsync(int id);
        IEnumerable<TodoModel> GetAll();
        IAsyncEnumerable<TodoModel> GetAllAsync();
        int? Insert(TodoModel model);
        bool Update(TodoModel model);
        bool ToggleComplete(int id);
        bool Delete(int id);
        bool DeleteAll();
    }
}