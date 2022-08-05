using Togo.Infrastructure.Identities.Dtos;

namespace Togo.Infrastructure.Identities;

public interface IUserService
{
    Task SeedAdminUserAsync();
    
    Task<UserDto> CreateAsync(CreateUserDto input);

    Task<List<UserDto>> GetAllAsync();

    Task<UserDto> GetByUserNameAsync(string userName);

    Task<bool> AuthenticateAsync(LoginDto input);
}
