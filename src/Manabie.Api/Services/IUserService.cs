

using Manabie.Api.Models;

using Manabie.Api.Entities;

namespace Manabie.Api.Services;

public interface IUserService
{
    User GetById(int userId);
    AuthenticateResponse Authenticate(AuthenticateRequest model);
    IEnumerable<User> GetAll();
}
