using Todo.Contract;
using Todo.Domain.Interfaces;
using Todo.Storage.Contract;
using Todo.Storage.Contract.Interfaces;

namespace Todo.Domain;

public class TodoManagementService : ITodoManagementService
{
    private readonly ITodoRepository _repository;

    public TodoManagementService(ITodoRepository repository)
    {
        _repository = repository ?? throw new ArgumentNullException(nameof(repository));
    }

    public async Task<Todo.Domain.Models.Todo> AddAsync(CreateTodoResource resource)
    {
        if (resource == null)
        {
            throw new ArgumentNullException("Resource must not be null.", nameof(resource));
        }

        var request = TranslateTodoRequest(resource);
        var result = await _repository.AddAsync(request);
        var todo = TranslateTodoResponse(result);

        return todo;
    }

    public async Task<long> DeleteAsync(long id)
    {
        if (id < 1)
        {
            return 0;
        }

        return await _repository.DeleteAsync(id);
    }

    // This method translates the client request to a Todo database request.
    private TodoRequest TranslateTodoRequest(CreateTodoResource resource)
    {
        DateTime dateNow = DateTime.UtcNow;

        return new TodoRequest
        {
            Name = resource.Name,
            Description = resource.Description,
            DateCreatedUTC = dateNow,
            DateModifiedUTC = dateNow,
            UserId = resource.UserId
        };
    }

    // This method translates the response from the database to the domain model Todo.
    private Todo.Domain.Models.Todo TranslateTodoResponse(TodoResponse response)
    {
        return new Todo.Domain.Models.Todo(
            response.Id,
            response.Name,
            response.Description,
            response.DateCreatedUTC,
            response.DateModifiedUTC,
            response.UserId
        );
    }
}