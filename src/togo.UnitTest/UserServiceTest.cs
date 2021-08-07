using NUnit.Framework;
using System;
using togo.Service;
using togo.Service.Errors;
using togo.Service.Interface;

namespace togo.UnitTest
{
    [TestFixture]
    public class UserServiceTest
    {
        private IUserService _userService;

        [SetUp]
        public void SetUp()
        {
            _userService = new UserService(TestHelper.GetMockTogoContext());
        }

        [Test]
        public void Login_ReturnToken()
        {
            var result = _userService.Login("firstUser", "example").Result;
            Assert.IsNotNull(result);
        }

        [Test]
        public void Login_ThrowException()
        {
            void _login()
            {
                try
                {
                    _userService.Login("firstUser", "examp").Wait();
                }
                catch (Exception e)
                {
                    throw e.InnerException;
                }

            };

            Assert.Throws(Is.TypeOf<RestException>(), _login);
        }
    }
}
