using Microsoft.Extensions.DependencyInjection;
using Services.Interfaces;
using Services.Models;
using Xunit;

namespace IntergationTest
{
    [Collection("Sequential")]
    public class UserServiceTest : BaseService_Test
    {
        private IUserService _userService;
        public UserServiceTest() : base()
        {
            _userService = ServiceProvider.GetRequiredService<IUserService>();
        }

        [Fact]
        public void Should_Register_Success()
        {
            var userRegister = new UserRegister()
            {
                Username = "Admin",
                Password = "123456789",
                TaskPerDay = 10
            };

            var result = _userService.Register(userRegister);

            Assert.True(result);
        }

        //[Fact]
        //public void Should_Register_Exception()
        //{
        //    var userRegister = new UserRegister()
        //    {
        //        Username = "TestUser1",
        //        Password = "123456789",
        //        TaskPerDay = 10
        //    };

        //    var result = _userService.Register(userRegister);

        //    Assert.False(result);
        //}

        [Fact]
        public void Should_Authenticate_Success()
        {
            var userLogin = new UserLogin()
            {
                Username = "TestUser1",
                Password = "123456789"
            };

            var result = _userService.Authenticate(userLogin);

            Assert.Equal("TestUser1", result.Username);
            Assert.False(result.NotFound);
        }

        [Fact]
        public void Should_Authenticate_UserOrPasswordIncorrect()
        {
            var userLogin = new UserLogin()
            {
                Username = "Member",
                Password = "123456789"
            };

            var result = _userService.Authenticate(userLogin);

            Assert.Null(result.ID);
            Assert.Null(result.Username);
            Assert.True(result.NotFound);
        }

        [Fact]
        public void Should_Authenticate_Exception()
        {
            var result = _userService.Authenticate(null);

            Assert.Null(result);
        }
    }
}
