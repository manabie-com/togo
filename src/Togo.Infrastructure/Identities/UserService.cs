using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Togo.Infrastructure.Identities.Dtos;

namespace Togo.Infrastructure.Identities;

public class UserService : IUserService
{
    private const string AdminUserName = "admin";
    private const string AdminPassword = "Abcd@1234";

    private readonly UserManager<AppUser> _userManager;
    private readonly RoleManager<IdentityRole> _roleManager;

    public UserService(UserManager<AppUser> userManager, RoleManager<IdentityRole> roleManager)
    {
        _userManager = userManager;
        _roleManager = roleManager;
    }

    public async Task SeedAdminUserAsync()
    {
        var role = await _roleManager.FindByNameAsync(Roles.Admin);
        if (role is null)
        {
            await _roleManager.CreateAsync(new IdentityRole(Roles.Admin));
            await _roleManager.FindByNameAsync(Roles.Admin);
        }
        
        var adminUser = await _userManager.FindByNameAsync(AdminUserName);
        if (adminUser is null)
        {
            var user = AppUser.Create(AdminUserName, int.MaxValue); 
            await _userManager.CreateAsync(user, AdminPassword);
            await _userManager.AddToRoleAsync(user, Roles.Admin);
        }
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
