using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using TODO.Repositories.Data;
using TODO.Repositories.Data.DBModels;
using TODO.Repositories.Interfaces;

namespace TODO.Repositories.Repositories
{
    public class UserRepository : IUserRepository
    {
        private readonly TodoContext _context;

        public UserRepository(TodoContext context)
        {
            _context = context ?? throw new ArgumentNullException(nameof(TodoContext));
        }

        public async Task<IEnumerable<User>> GetUsers(int userId)
        {
            try
            {
                var result = await _context.User
                                .Include(x => x.UserTodoConfig)
                                .Include(x => x.Todos)
                                .Where(u => userId > 0 ? u.UserId == userId : true)
                                .AsNoTracking()
                                .ToListAsync();

                return result;
            }
            catch (Exception e)
            {

                throw;
            }
        }
    }
}
