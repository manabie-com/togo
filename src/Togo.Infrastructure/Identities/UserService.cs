using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging;
using Microsoft.IdentityModel.Tokens;
using Togo.Core;
using Togo.Core.Exceptions;
using Togo.Infrastructure.Identities.Dtos;
using Togo.Infrastructure.Identities.Exceptions;

namespace Togo.Infrastructure.Identities;

public class UserService : IUserService
{
    private const string AdminUserName = "admin";
    private const string AdminPassword = "Abcd@1234";

    private readonly UserManager<AppUser> _userManager;
    private readonly RoleManager<IdentityRole> _roleManager;
    private readonly TogoAppSettings _togoAppSettings;
    private readonly ILogger<UserService> _logger;

    public UserService(
        UserManager<AppUser> userManager, 
        RoleManager<IdentityRole> roleManager,
        ILogger<UserService> logger, 
        TogoAppSettings togoAppSettings)
    {
        _userManager = userManager;
        _roleManager = roleManager;
        _logger = logger;
        _togoAppSettings = togoAppSettings;
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

        _logger.LogError("User creation failed");
        throw new UserCreationFailedException(result);
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

    public async Task<LoginResponseDto> AuthenticateAsync(LoginDto input) 
    {
        var user = await _userManager.FindByNameAsync(input.UserName);

        if (user == null)
        {
            _logger.LogWarning("User {UserName} is not found", input.UserName);
            throw new InvalidLoginException();
        }

        if (!await _userManager.CheckPasswordAsync(user, input.Password))
        {
            _logger.LogWarning("User {UserName} logged in with wrong password", input.UserName);
            throw new InvalidLoginException();
        }

        var sessionId = Guid.NewGuid().ToString();
        var roles = await _userManager.GetRolesAsync(user);

        return new LoginResponseDto(input.UserName, IssueToken(sessionId, user.Id, roles, user.MaxTasksPerDay));
    }
    
    private string IssueToken(
        string sessionId,
        string userId,
        IEnumerable<string> roleIds,
        int maxTasksPerDay)
    {
        var securityKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_togoAppSettings.JwtBearer.SecurityKey));
        var signingCredentials = new SigningCredentials(securityKey, SecurityAlgorithms.HmacSha256Signature);

        var now = DateTime.UtcNow;

        var claims = new List<Claim>
        {
            new(JwtRegisteredClaimNames.Sid, sessionId),
            new(JwtRegisteredClaimNames.Sub, userId),
            new(TogoCustomClaims.MaxTasksPerDay, maxTasksPerDay.ToString())
        };

        claims.AddRange(roleIds.Select(roleId => new Claim(ClaimTypes.Role, roleId)));

        var jwtSecurityToken = new JwtSecurityToken(
            _togoAppSettings.JwtBearer.Issuer,
            _togoAppSettings.JwtBearer.Audience,
            claims,
            now,
            now.Add(TimeSpan.FromHours(8)),
            signingCredentials);

        return new JwtSecurityTokenHandler().WriteToken(jwtSecurityToken);
    }
}
