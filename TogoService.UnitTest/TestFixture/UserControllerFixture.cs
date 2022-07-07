using System;
using Autofac;
using AutoMapper;
using Microsoft.Extensions.Logging;
using Moq;
using TogoService.API.Controller;
using TogoService.API.Infrastructure.Mapper;
using TogoService.API.Infrastructure.Repository;
using TogoService.API.Model.Interface;

namespace TogoService.UnitTest.TestFixture
{
    public class UserControllerFixture : IDisposable
    {
        public IContainer Container { get; private set; }
        public Mock<IUnitOfWork> MockUnitOfWork { get; set; }
        public Mock<BaseRepository<API.Model.User>> MockGenericUserRepository { get; set; }
        public Mock<ITodoTaskRepository> MockTodoTaskRepository { get; set; }

        public UserControllerFixture()
        {
            // Register all DI need for unit test.
            var builder = new ContainerBuilder();

            // Mapper
            builder.Register(context => new MapperConfiguration(cfg =>
            {
                cfg.AddProfile(new AutoMapperProfile());
            }));
            builder.Register(c =>
            {
                var context = c.Resolve<IComponentContext>();
                var config = context.Resolve<MapperConfiguration>();
                return config.CreateMapper(context.Resolve);
            }).As<IMapper>().SingleInstance();

            var logger = new Mock<ILogger<UserController>>();
            MockUnitOfWork = new Mock<IUnitOfWork>();
            MockGenericUserRepository = new Mock<BaseRepository<API.Model.User>>();
            MockTodoTaskRepository = new Mock<ITodoTaskRepository>();

            builder.RegisterType<UserController>().As<UserController>()
                .WithParameter(new TypedParameter(typeof(ILogger<UserController>), logger.Object))
                .WithParameter(new TypedParameter(typeof(IUnitOfWork), MockUnitOfWork.Object))
                .WithParameter(new TypedParameter(typeof(ITodoTaskRepository), MockTodoTaskRepository.Object));

            //Build container
            Container = builder.Build();
        }

        public void Dispose()
        {
            Container.Dispose();
        }
    }
}
