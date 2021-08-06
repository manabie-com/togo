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
                    Password = "example",
                    MaxTodo = 5,
                };

                await context.AddAsync(user);
                await context.SaveChangesAsync();
            }
        }
    }
}
