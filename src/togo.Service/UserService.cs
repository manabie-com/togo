using System.Threading.Tasks;
using togo.Service.Context;
using togo.Service.Interface;
using System.Linq;
using Microsoft.EntityFrameworkCore;
using togo.Service.Errors;
using System.Net;
using togo.Service.Helper;

namespace togo.Service
{
    public class UserService : IUserService
    {
        private readonly TogoContext _context;

        public UserService(TogoContext context)
        {
            _context = context;
        }

        public async Task<string> Login(string userId, string password)
        {
            await ValidateUser(userId, password);
            return JwtGeneratorHelper.CreateToken(userId);
        }

        private async System.Threading.Tasks.Task ValidateUser(string userId, string password)
        {
            var query = from u in _context.Users
                        where u.Id == userId
                        select u;
            var user = await query.FirstOrDefaultAsync();

            var isValid = SercurityHelper.ComparePassword(password, user?.PasswordSalt, user?.PasswordHash);
            if (!isValid)
            {
                throw new RestException(HttpStatusCode.Unauthorized);
            }
        }
    }
}
