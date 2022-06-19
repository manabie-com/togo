using Microsoft.EntityFrameworkCore;
using Todo.Application.Interfaces;
using Todo.Domain.Entities;

namespace Todo.Infrastructure
{
    public class ApplicationDbSeeder
    {
        private readonly IApplicationDbContext _dbContext;
        public ApplicationDbSeeder(IApplicationDbContext dbContext)
        {
            _dbContext = dbContext;
        }

        public void EnsureMigrate()
        {
            _dbContext.Database.Migrate();
        }
        public void EnsureData()
        {
            AddDefaultUser();
        }
        private void AddDefaultUser() { 
            var users = _dbContext.Users;
            if (!users.Any())
            {
                var alice = new User
                {
                    Name = "Alice",
                    LimitTask = 20,
                    CreatedAt = DateTime.Now,
                    CreatedBy = "System"
                };
                _dbContext.Users.Add(alice);

                var bob = new User
                {
                    Name = "Bob",
                    LimitTask = 30,
                    CreatedAt = DateTime.Now,
                    CreatedBy = "System"
                };
                _dbContext.Users.Add(bob);

                _dbContext.SaveChanges();
            }

        }
    }
}
