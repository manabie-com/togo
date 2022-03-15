using Todo.Contract;
using Todo.Domain.Interfaces;
using Todo.Domain.Models;
using Todo.Storage.Contract;
using Todo.Storage.Contract.Interfaces;

namespace Todo.Domain;

public class UserManagementService : IUserManagementService
{
    private readonly IUserRepository _repository;

    private readonly ISerializer _serializer;

    public UserManagementService(IUserRepository repository, ISerializer serializer)
    {
        _repository = repository ?? throw new ArgumentNullException(nameof(repository));
        _serializer = serializer ?? throw new ArgumentNullException(nameof(serializer));
    }

    public async Task<Todo.Domain.Models.User> AddAsync(CreateUserResource resource)
    {
        if (resource == null)
        {
            throw new ArgumentNullException("Resource must not be null.", nameof(resource));
        }

        var request = TranslateUserRequest(resource);
        var result = await _repository.AddAsync(request);
        var user = TranslateUserResponse(result);

        return user;
    }

    public async Task<User> GetAsync(long id)
    {
        if (id < 1)
        {
            throw new ArgumentOutOfRangeException("Id must be greater than 0.", nameof(id));
        }

        var result = await _repository.GetAsync(id);
        var user = TranslateUserResponse(result);

        return user;
    }

    public async Task<long> DeleteAsync(long id)
    {
        if (id < 1)
        {
            return 0;
        }

        return await _repository.DeleteAsync(id);
    }

    // This method translates the client request to a user request for the database.
    private UserRequest TranslateUserRequest(CreateUserResource resource)
    {
        return new UserRequest
        {
            FirstName = resource.FirstName,
            LastName = resource.LastName,
            Todos = "",
            DailyTaskLimit = resource.DailyTaskLimit
        };
    }

    // This method translates the response from the database back to the domain model User.
    private User TranslateUserResponse(UserResponse response)
    {
        var user = new User(
            response.Id,
            response.FirstName,
            response.LastName,
            response.DailyTaskLimit
        );

        if (!string.IsNullOrEmpty(response.Todos))
        {
            user.Todos = _serializer.Deserialize<IEnumerable<Todo.Domain.Models.Todo>>(response.Todos);
        }

        return user;
    }
}