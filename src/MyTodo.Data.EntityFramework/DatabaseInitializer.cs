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
                await roleManager.CreateAsync(new AppRole()
                {
                    Name = "Administration",
                    NormalizedName = "Administration",
                });
                await roleManager.CreateAsync(new AppRole()
                {
                    Name = "Staff",
                    NormalizedName = "Staff",
                });
                context.SaveChanges();
            }

            if (!userManager.Users.Any())
            {

                var user01 = new AppUser()
                {
                    UserName = "user01@mytodo.com",
                    Email = "user01@mytodo.com",
                    PhoneNumber = "0989774788",
                    TaskCount = 0,
                    TaskLimit = 100,

                };

                var result = await userManager.CreateAsync(user01, "123456");
                if (result.Succeeded)
                    await userManager.AddToRoleAsync(user01, "Administration");


                var user02 = new AppUser()
                {
                    UserName = "user02@gmail.com",
                    Email = "user02@gmail.com",
                    TaskCount = 0,
                    TaskLimit = 5

                };

                result = await userManager.CreateAsync(user02, "123456");
                if (result.Succeeded)
                {
                    await userManager.AddToRoleAsync(user02, "Staff");

                }

                var user03 = new AppUser()
                {
                    UserName = "user03@gmail.com",
                    Email = "user03@gmail.com",
                    TaskCount = 0,
                    TaskLimit = 5

                };
                result = await userManager.CreateAsync(user03, "123456");
                if (result.Succeeded)
                {
                    await userManager.AddToRoleAsync(user03, "Staff");

                }

                //

            }
            if (!context.TodoItems.Any())
            {
                var todoItems = new List<TodoItem>(){
                            new TodoItem(){ Title="Task 1", Description="Task 1", Priority=1,Status=Enums.TodoItemStatus.New },
                            new TodoItem(){ Title="Task 2", Description="Task 2", Priority=2,Status=Enums.TodoItemStatus.New },
                            new TodoItem(){ Title="Task 3", Description="Task 3", Priority=3,Status=Enums.TodoItemStatus.New },
                            new TodoItem(){ Title="Task 4", Description="Task 4", Priority=1,Status=Enums.TodoItemStatus.New },
                            new TodoItem(){ Title="Task 5", Description="Task 5", Priority=2,Status=Enums.TodoItemStatus.New },
                            new TodoItem(){ Title="Task 6", Description="Task 6", Priority=3,Status=Enums.TodoItemStatus.New },
                        };
                context.TodoItems.AddRange(todoItems);
            }
            context.SaveChanges();
        }
    }
}
