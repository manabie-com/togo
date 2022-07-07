using Autofac;
using AutoMapper;
using TogoService.API.Controller;
using TogoService.API.Model;
using TogoService.IntegrationTest.TestFixture;
using Xunit;

namespace TogoService.IntegrationTest.Tests
{

    public partial class UserControllerTest : IClassFixture<UserControllerFixture>
    {
        private UserController _userController;
        private IMapper _mapper;
        private User _userWith0MaxDailyTasks;
        private User _userWith10MaxDailyTasks;

        public UserControllerTest(UserControllerFixture fixture)
        {
            using (var scope = fixture.Container.BeginLifetimeScope())
            {
                _userController = scope.Resolve<UserController>();
                _mapper = scope.Resolve<IMapper>();
            }

            _userWith0MaxDailyTasks = fixture.UserWith0MaxDailyTasks;
            _userWith10MaxDailyTasks = fixture.UserWith10MaxDailyTasks;
        }
    }
}