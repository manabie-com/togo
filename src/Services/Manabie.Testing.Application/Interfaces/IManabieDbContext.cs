using Manabie.Testing.Domain.Entities;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Testing.Application.Interfaces
{
    public  interface IManabieDbContext
    {
        public DbSet<Todo> Todos { get; }
        public DbSet<AppSetting> AppSettings { get; }
        public DbSet<UserLimit> UserLimits { get; }
        Task<int> SaveChangesAsync(CancellationToken cancellationToken = default);
    }
}
