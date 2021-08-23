using Application.Interfaces.Repositories;
using Domain.Entities;
using Domain.Settings;
using Infrastructure.Persistence.Contexts;
using Microsoft.EntityFrameworkCore;
using System;

namespace Infrastructure.Persistence.Seeds
{
    public static class DefaultBasicUser
    {
        public static async System.Threading.Tasks.Task SeedAsync(ApplicationDbContext applicationDbContext, IUserServiceAsync userService)
        {
            var password = "example";
            var hassPassword = userService.Encrypt(password);
            //Seed Default User
            var defaultUser = new User
            {
                Email = "firstUser@gmail.com",
                Id = "firstUser",
                Password = hassPassword,
                CreatedDate = DateTime.UtcNow,
                CreatedBy = "system"
            };
            var userExist = await applicationDbContext.Users.FirstOrDefaultAsync(_ => _.Email == defaultUser.Email);
            if (userExist == null)
            {
                await applicationDbContext.Users.AddAsync(defaultUser);
                await applicationDbContext.SaveChangesAsync();
            }
        }
    }
}
