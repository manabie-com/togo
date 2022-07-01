using Manabie.Testing.Application.Interfaces;
using Manabie.Testing.Domain.Entities;
using Microsoft.EntityFrameworkCore;

namespace Manabie.Testing.Infrastructure.Persistence
{
    public class ManabieDbContext : DbContext, IManabieDbContext
    {
        public ManabieDbContext(DbContextOptions<ManabieDbContext> options) : base(options)
        {
        }

        protected override void OnModelCreating(ModelBuilder builder)
        {

            base.OnModelCreating(builder);
        }

        public DbSet<Todo> Todos => Set<Todo>();

        public DbSet<AppSetting> AppSettings => Set<AppSetting>();
        public DbSet<UserLimit> UserLimits => Set<UserLimit>();
    }
}
