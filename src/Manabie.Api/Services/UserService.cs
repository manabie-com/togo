

using Manabie.Api.Insfrastructures;
using Manabie.Api.Models;
using Manabie.Api.Utilities;

using Manabie.Api.Entities;

using Microsoft.Extensions.Options;
using static Manabie.Api.Utilities.Utility;
using Microsoft.AspNetCore.Identity;

namespace Manabie.Api.Services;

public class UserService : IUserService
{
    private readonly ManaDbContext _context;
    private readonly AppSettings _appSettings;

    public UserService(ManaDbContext context, IOptions<AppSettings> appSettings)
    {
        _context = context;
        _appSettings = appSettings.Value;
    }

    public AuthenticateResponse Authenticate(AuthenticateRequest model)
    {
        var password = HashSh256(model.Password, _appSettings.Secret);
        var user = _context.Users.SingleOrDefault(x => x.Username == model.Username && x.Password == password);

        // return null if user not found
        if (user == null) return null;

        // authentication successful so generate jwt token
        var token = GenerateJwtToken(user, _appSettings.Secret);

        return new AuthenticateResponse(user, token);
    }

    public IEnumerable<User> GetAll()
    {
        return _context.Users.ToList();
    }

    public User GetById(int userId)
    {
        return _context.Users.Find(userId);
    }
}
