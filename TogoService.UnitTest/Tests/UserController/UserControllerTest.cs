using Autofac;
using AutoMapper;
using Moq;
using TogoService.API.Controller;
using TogoService.API.Infrastructure.Repository;
using TogoService.API.Model.Interface;
using TogoService.UnitTest.TestFixture;
using Xunit;

namespace TogoService.UnitTest.Tests
{

    public partial class UserControllerTest : IClassFixture<UserControllerFixture>
    {
        private UserController _userController;
        private IMapper _mapper;
        private Mock<IUnitOfWork> _mockUnitOfWork { get; set; }
        private Mock<BaseRepository<API.Model.User>> _mockGenericUserRepository { get; set; }
        private Mock<ITodoTaskRepository> _mockTodoTaskRepository { get; set; }

        public UserControllerTest(UserControllerFixture fixture)
        {
            using (var scope = fixture.Container.BeginLifetimeScope())
            {
                _userController = scope.Resolve<UserController>();
                _mapper = scope.Resolve<IMapper>();
            }

            _mockUnitOfWork = fixture.MockUnitOfWork;
            _mockGenericUserRepository = fixture.MockGenericUserRepository;
            _mockUnitOfWork.Setup(q => q.GenericRepository<API.Model.User>()).Returns(_mockGenericUserRepository.Object);
            _mockTodoTaskRepository = fixture.MockTodoTaskRepository;
        }
    }
}