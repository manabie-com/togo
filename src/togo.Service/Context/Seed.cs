using System.Linq;
using Future = System.Threading.Tasks.Task;

namespace togo.Service.Context
{
    public class Seed
    {
        public static async Future SeedData(TogoContext context)
        {
            if (!context.Users.Any())
            {
                var user = new User
                {
                    Id = "firstUser",
                    PasswordSalt = "q+/pqaqsAslRAgZOK4b0ug==",
                    PasswordHash = "KxcQVXazTAsDE6+oIXW6aAeqzsLglHImlaleuVAvyMs=",
                    MaxTodo = 5,
                };

                await context.AddAsync(user);
                await context.SaveChangesAsync();
            }
        }
    }
}
