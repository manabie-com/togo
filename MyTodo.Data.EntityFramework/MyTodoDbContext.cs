using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore;
using MyTodo.Data.Entities;
using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Data.EntityFramework
{
    public class MyTodoDbContext : IdentityDbContext<AppUser, AppRole, Guid>
    {
        public MyTodoDbContext()
        {

        }
        public MyTodoDbContext(DbContextOptions options) : base(options)
        {

        }
        public DbSet<AppUser> AppUsers { get; set; }
        public DbSet<AppRole> AppRoles { get; set; }
        public DbSet<TodoItem> TodoItem { get; set; }
        public DbSet<Assignment> Assignments { get; set; }
        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            optionsBuilder.UseSqlServer(@"Server=LA-TUAN;Database=MyTodo;Trusted_Connection=True;");
        }

    }
}
