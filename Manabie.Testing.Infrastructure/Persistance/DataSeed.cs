using Manabie.Testing.Application.Interfaces;
using Manabie.Testing.Domain.Entities;
using Manabie.Testing.Infrastructure.Persistence;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Testing.Infrastructure.Persistance
{
    public class DataSeed
    {
        private readonly ILogger<DataSeed> _logger;
        private readonly ManabieDbContext _context;
        public DataSeed(ILogger<DataSeed> logger, ManabieDbContext context)
        {
            _logger = logger;
            _context = context;
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

            var settings = new List<AppSetting>()
            {
                new AppSetting()
                {
                    Key = "User", Value = "15"
                },
                new AppSetting()
                {
                    Key = "Administrator", Value = "10"
                }
            };
            await _context.AppSettings.AddRangeAsync(settings);

            await _context.SaveChangesAsync();
        }
    }
}
