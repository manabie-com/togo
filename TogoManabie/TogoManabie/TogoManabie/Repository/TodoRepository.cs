using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using TogoManabie.Configuration;
using TogoManabie.Interfaces;
using TogoManabie.Models;

namespace TogoManabie.Repository
{
    public class TodoRepository : BaseRepository<Tasks>
    {
        private readonly EntityFrameworkSqlServerContext _dbContext;
        public TodoRepository(EntityFrameworkSqlServerContext dbContext) : base(dbContext)
        {
            
        }
        public async Task<List<Tasks>> GetAllByUserId(int userId , DateTime now)
        {
            return await _dbContext.Set<Tasks>()
                .AsNoTracking()
                .Where(s => s.user_id == userId && s.created_date == now)
                .ToListAsync();
        }
    }
}
