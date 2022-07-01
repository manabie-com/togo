using Manabie.BasicIdentityServer.Infrastructure.Identity;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging;

namespace Manabie.BasicIdentityServer.Infrastructure.Persistence
{
    public class DataSeed
    {
        private readonly ILogger<DataSeed> _logger;
        private readonly ApplicationDbContext _context;
        private readonly UserManager<ApplicationUser> _userManager;
        private readonly RoleManager<IdentityRole> _roleManager;

        public DataSeed(ILogger<DataSeed> logger, ApplicationDbContext context, UserManager<ApplicationUser> userManager, RoleManager<IdentityRole> roleManager)
        {
            _logger = logger;
            _context = context;
            _userManager = userManager;
            _roleManager = roleManager;
        }

        public async Task InitialiseAsync()
        {
            try
            {
                if (_context.Database.IsRelational())
                {
                    await _context.Database.MigrateAsync();
                }
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, "An error occurred while initialising the database.");
                throw;
            }
        }

        public async Task SeedAsync()
        {
            try
            {
                await TrySeedAsync();
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, "An error occurred while seeding the database.");
                throw;
            }
        }

        public async Task TrySeedAsync()
        {
            // Default roles
            var administratorRole = new IdentityRole("Administrator");
            var userRole = new IdentityRole("User");

            if (!_roleManager.Roles.Any())
            {
                await _roleManager.CreateAsync(administratorRole);
                await _roleManager.CreateAsync(userRole);
            }

            await _context.SaveChangesAsync();

            // Default users
            var administrator = new ApplicationUser { UserName = "administrator@localhost", Email = "administrator@localhost" };
            var user = new ApplicationUser { UserName = "user@localhost", Email = "user@localhost" };

            if (!_userManager.Users.Any())
            {
                await _userManager.CreateAsync(administrator, "Administrator1!");
                await _userManager.AddToRolesAsync(administrator, new[] { administratorRole.Name });

                await _userManager.CreateAsync(user, "User1!");
                await _userManager.AddToRolesAsync(user, new[] { userRole.Name });
            }
        }
    }
}
