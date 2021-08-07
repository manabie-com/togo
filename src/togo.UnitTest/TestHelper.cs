using Microsoft.EntityFrameworkCore;
using Moq;
using togo.Service;
using togo.Service.Context;
using togo.Service.Interface;

namespace togo.UnitTest
{
    public class TestHelper
    {
        private static TogoContext _context;
        public static TogoContext GetMockTogoContext()
        {
            if (_context is null)
            {
                var options = new DbContextOptionsBuilder<TogoContext>()
                .UseInMemoryDatabase(databaseName: "TogoDb")
                .Options;

                var context = new TogoContext(options);
                context.Users.Add(new User
                {
                    Id = "firstUser",
                    PasswordSalt = "q+/pqaqsAslRAgZOK4b0ug==",
                    PasswordHash = "KxcQVXazTAsDE6+oIXW6aAeqzsLglHImlaleuVAvyMs=",
                    MaxTodo = 5,
                });
                context.SaveChanges();
                _context = context;
            }
            return _context;
        }

        private static ICurrentHttpContext _CurrentHttpContext;
        public static ICurrentHttpContext GetMockCurrentHttpContext()
        {
            if (_CurrentHttpContext is null)
            {
                var mockHttpCurrentContext = new Mock<ICurrentHttpContext>();
                mockHttpCurrentContext.Setup(m => m.GetCurrentUserId()).Returns("firstUser");
                _CurrentHttpContext = mockHttpCurrentContext.Object;
            }
            return _CurrentHttpContext;
        }
    }
}
