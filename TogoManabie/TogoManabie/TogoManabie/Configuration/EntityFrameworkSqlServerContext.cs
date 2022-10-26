using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using TogoManabie.Models;

namespace TogoManabie.Configuration
{
    public class EntityFrameworkSqlServerContext : DbContext
    {
        public EntityFrameworkSqlServerContext(DbContextOptions<EntityFrameworkSqlServerContext> options) : base(options)
        {
            this.Database.EnsureCreated();
        }

        public DbSet<Models.Task> Tasks
        {
            get;
            set;
        }
        public DbSet<Models.User> Users
        {
            get;
            set;
        }
        
    }
}
