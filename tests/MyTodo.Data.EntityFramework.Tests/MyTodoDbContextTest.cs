using System;
using System.Collections.Generic;
using System.Text;
using Xunit;

namespace MyTodo.Data.EntityFramework.Tests
{
    public class MyTodoDbContextTest
    {
        [Fact]
        public void Constructor_CreateInMemoryDb_Success()
        {
            var context = ContextFactory.Create();
            Assert.True(context.Database.EnsureCreated());
        }
    }
}
