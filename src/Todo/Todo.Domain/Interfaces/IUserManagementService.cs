using Todo.Contract;
using Todo.Domain.Models;

namespace Todo.Domain.Interfaces;

public interface IUserManagementService
{
    Task<User> AddAsync(CreateUserResource resource);

    Task<User> GetAsync(long id);

    Task<long> DeleteAsync(long id);
}