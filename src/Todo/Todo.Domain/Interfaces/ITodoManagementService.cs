using Todo.Contract;

namespace Todo.Domain.Interfaces;

public interface ITodoManagementService
{
    Task<Todo.Domain.Models.Todo> AddAsync(CreateTodoResource resource);

    Task<long> DeleteAsync(long id);
}