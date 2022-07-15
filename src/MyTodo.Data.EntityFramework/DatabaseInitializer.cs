using Microsoft.AspNetCore.Identity;
using MyTodo.Data.Entities;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace MyTodo.Data.EntityFramework
{
    public class DatabaseInitializer
    {
        private readonly MyTodoDbContext context;
        private readonly UserManager<AppUser> userManager;
        private readonly RoleManager<AppRole> roleManager;
        public DatabaseInitializer(MyTodoDbContext context, UserManager<AppUser> userManager, RoleManager<AppRole> roleManager)
        {
            this.context = context;
            this.userManager = userManager;
            this.roleManager = roleManager;
        }
        public async Task Seed()
        {
            if (!roleManager.Roles.Any())
            {
                await roleManager.CreateAsync(new AppRole() { Name = "Admin" });
                await roleManager.CreateAsync(new AppRole() { Name = "User" });
                context.SaveChanges();
            }

            if (!userManager.Users.Any())
            {
                var adminAccount = new AppUser()
                {
                    UserName = "admin@mytodo.com",
                    Email = "admin@mytodo.com",
                    PhoneNumber = "0989774722",
                    TaskCount = 0,
                    TaskLimit = 100,

                };

                var result = await userManager.CreateAsync(adminAccount, "123456");
                if (result.Succeeded)
                {
                    await userManager.AddToRoleAsync(adminAccount, "Admin");

                }

                var userAccount = new AppUser()
                {
                    UserName = "user@mytodo.com",
                    Email = "user@mytodo.com",
                    PhoneNumber = "0989774723",
                    TaskCount = 0,
                    TaskLimit = 2,

                };

                result = await userManager.CreateAsync(userAccount, "123456");
                if (result.Succeeded)
                {
                    await userManager.AddToRoleAsync(userAccount, "User");

                }
                //

            }
            if (!context.TodoItems.Any())
            {
                var todoItems = new List<TodoItem>(){
                            new TodoItem(){ Title="Task 1", Description="Task 1", Priority=1,Status=Enums.TodoItemStatus.New },
                            new TodoItem(){ Title="Task 2", Description="Task 2", Priority=2,Status=Enums.TodoItemStatus.New },
                        };
                context.TodoItems.AddRange(todoItems);
            }
            context.SaveChanges();
        }
    }
}
