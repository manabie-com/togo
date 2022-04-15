using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using TogoManabie.Configuration;
using TogoManabie.Models;

namespace TogoManabie.Repository
{
    public class UserRepository : BaseRepository<User>
    {
        public UserRepository(EntityFrameworkSqlServerContext dbContext) : base(dbContext)
        {

        }
    }
}
