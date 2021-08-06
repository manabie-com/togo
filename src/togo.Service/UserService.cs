using System;
using System.Threading.Tasks;
using togo.Service.Context;
using togo.Service.Interface;
using System.Linq;
using Microsoft.EntityFrameworkCore;

namespace togo.Service
{
    public class UserService : IUserService
    {
        private readonly TogoContext _context;
        private readonly IJwtGenerator _jwtGenerator;

        public UserService(
              TogoContext context
            , IJwtGenerator jwtGenerator)
        {
            _context = context;
            _jwtGenerator = jwtGenerator;
        }

        public async Task<(bool, string)> Login(string userId, string password)
        {
            var isValid = await ValidateUser(userId, password);
            if (!isValid)
            {
                return (isValid, string.Empty);
            }
            return (isValid, _jwtGenerator.CreateToken(userId));
        }

        private async Task<bool> ValidateUser(string userId, string password)
        {
            var query = from u in _context.Users
                        where u.Id == userId && u.Password == password
                        select u;

            return await query.AnyAsync();
        }
    }
}
