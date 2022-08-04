using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Togo.Infrastructure.Identities.Dtos;

namespace Togo.Infrastructure.Identities;

public class UserService : IUserService
{
    private readonly UserManager<AppUser> _userManager;

    public UserService(UserManager<AppUser> userManager)
    {
        _userManager = userManager;
    }

    public async Task<UserDto> CreateAsync(CreateUserDto input)
    {
        var user = AppUser.Create(input.UserName, input.MaxTasksPerDay);
        var result = await _userManager.CreateAsync(user, input.Password);

        if (result.Succeeded)
        {
            return new UserDto(user);
        }

        throw new Exception("User creation failed");
    }

    public async Task<List<UserDto>> GetAllAsync()
    {
        return await _userManager.Users
            .Select(user => new UserDto(user))
            .ToListAsync();
    }

    public async Task<UserDto> GetByUserNameAsync(string userName)
    {
        var user = await _userManager.FindByNameAsync(userName);
        return new UserDto(user);
    }

    public async Task<bool> AuthenticateAsync(LoginDto input)
    {
        var user = await _userManager.FindByNameAsync(input.UserName);

        if (user == null)
        {
            throw new Exception("User not found");
        }

        return await _userManager.CheckPasswordAsync(user, input.Password);
    }
}