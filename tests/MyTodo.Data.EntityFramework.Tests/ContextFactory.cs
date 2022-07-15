using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Data.EntityFramework.Tests
{
    public class ContextFactory
    {
        public static MyTodoDbContext Create()
        {
            DbContextOptions<MyTodoDbContext> options = new DbContextOptionsBuilder<MyTodoDbContext>()
                .UseInMemoryDatabase(Guid.NewGuid().ToString()).Options;
            return new MyTodoDbContext(options);
        }
    }
}
