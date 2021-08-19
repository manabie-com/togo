using Microsoft.Extensions.Configuration;
using MockQueryable.Moq;
using Moq;
using System.Collections.Generic;
using WebApi.Models;
using WebApi.Services;
using Xunit;

namespace WebApi.Test
{
    public class UserServiceTest
    {
        private readonly UserService _userService;
        private readonly Mock<DemoDbContext> _mockDbContext;
        private readonly Dictionary<string, string> inMemorySettings = new Dictionary<string, string> {
            {"JWT:Secret", "ByYM000OLlMQG6VVVp1OH7Xzyr7gHuw1qvUC5dcGt3SNM"}
        };

        public UserServiceTest()
        {
            IConfiguration configuration = new ConfigurationBuilder()
                .AddInMemoryCollection(inMemorySettings)
                .Build();
            _mockDbContext = new Mock<DemoDbContext>();
            _userService = new UserService(_mockDbContext.Object, configuration);
        }

        [Theory]
        [InlineData("ee08f09c-319c-484c-936f-0c020e343bf5", "123456", "ee08f09c-319c-484c-936f-0c020e343bf5|2429507e-d057-44cb-9a6a-cce9447199a7|e75b9289-0ed6-45a2-96c6-b12848f958c6|bc347f5c-5c0b-4157-9d37-8645a3b11302|f2e120e9-1fd0-43a1-92f1-651a389a972a")]
        [InlineData("f2e120e9-1fd0-43a1-92f1-651a389a972a", "abc123", "e75b9289-0ed6-45a2-96c6-b12848f958c6|bc347f5c-5c0b-4157-9d37-8645a3b11302|f2e120e9-1fd0-43a1-92f1-651a389a972a")]
        public async System.Threading.Tasks.Task Login_UserNull(string userId, string password, string userIds)
        {
            // Setup
            var users = TestServiceUtilities.SetupDataForUserEntity(userIds, password);
            var usersMock = users.BuildMockDbSet();
            _mockDbContext.Setup(x => x.Users).Returns(usersMock.Object);

            var userLogin = await _userService.Login(System.Guid.Parse(userId), password);

            // Assert
            Assert.NotNull(userLogin);
        }

        [Theory]
        [InlineData("ee08f09c-319c-484c-936f-0c020e343bf5", "123456", "ee08f09c-319c-484c-936f-0c020e343bf5|2429507e-d057-44cb-9a6a-cce9447199a7|e75b9289-0ed6-45a2-96c6-b12848f958c6|bc347f5c-5c0b-4157-9d37-8645a3b11302|f2e120e9-1fd0-43a1-92f1-651a389a972a")]
        [InlineData("f2e120e9-1fd0-43a1-92f1-651a389a972a", "abc123", "e75b9289-0ed6-45a2-96c6-b12848f958c6|bc347f5c-5c0b-4157-9d37-8645a3b11302|f2e120e9-1fd0-43a1-92f1-651a389a972a")]
        public async System.Threading.Tasks.Task Login_UserHasData(string userId, string password, string userIds)
        {
            // Setup
            var users = TestServiceUtilities.SetupDataForUserEntity(userIds, password);
            var usersMock = users.BuildMockDbSet();
            _mockDbContext.Setup(x => x.Users).Returns(usersMock.Object);

            var userLogin = await _userService.Login(System.Guid.Parse(userId), password);

            // Assert
            Assert.NotNull(userLogin);
            Assert.NotEqual(System.Guid.Empty, userLogin.Id);
            Assert.Equal(userId, userLogin.Id.ToString());
            Assert.False(string.IsNullOrEmpty(userLogin.Token));
            // Check JWT token structure
            Assert.Equal(3, userLogin.Token.Split(".").Length);
        }
    }
}
