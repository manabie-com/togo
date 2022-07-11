using System;
using System.Threading.Tasks;
using Autofac;
using AutoMapper;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Diagnostics;
using Microsoft.Extensions.Logging;
using Moq;
using TogoService.API.Controller;
using TogoService.API.Infrastructure.Database;
using TogoService.API.Infrastructure.Mapper;
using TogoService.API.Infrastructure.Repository;
using TogoService.API.Model;
using TogoService.API.Model.Interface;
using TogoService.IntegrationTest.Helper;

namespace TogoService.IntegrationTest.TestFixture
{
    public class UserControllerFixture : IDisposable
    {
        public IContainer Container { get; private set; }
        public User UserWith0MaxDailyTasks { get; set; }
        public User UserWith10MaxDailyTasks { get; set; }

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

            // DbContext
            var contextOptions = new DbContextOptionsBuilder<TogoDbContext>()
                .UseLazyLoadingProxies()
                .ConfigureWarnings(b => b.Log((RelationalEventId.CommandCreated, LogLevel.Debug),
                    (RelationalEventId.CommandExecuted, LogLevel.Debug),
                    (RelationalEventId.TransactionCommitted, LogLevel.Debug)))
                .UseSqlite(@"Data Source=TogoService.db;")
                .Options;
            builder.RegisterType<TogoDbContext>()
                .WithParameter(new TypedParameter(typeof(DbContextOptions<TogoDbContext>), contextOptions));

            builder.RegisterType<UnitOfWork>().As<IUnitOfWork>();
            builder.RegisterType<TodoTaskRepository>().As<ITodoTaskRepository>();

            var logger = new Mock<ILogger<UserController>>();

            builder.RegisterType<UserController>().As<UserController>()
                .WithParameter(new TypedParameter(typeof(ILogger<UserController>), logger.Object))
                .SingleInstance();

            //Build container
            Container = builder.Build();

            Task task = InitData(new TogoDbContext(contextOptions));
            task.Wait();
        }

        public void Dispose()
        {
            Container.Dispose();
        }

        private async Task InitData(TogoDbContext dbContext)
        {
            IUnitOfWork unitOfWork = new UnitOfWork(dbContext);
            UserWith0MaxDailyTasks = FakeData.GenerateUser(0);
            UserWith10MaxDailyTasks = FakeData.GenerateUser(10);
            await unitOfWork.GenericRepository<User>().Add(UserWith0MaxDailyTasks);
            await unitOfWork.GenericRepository<User>().Add(UserWith10MaxDailyTasks);
            await unitOfWork.Save();
            dbContext.Dispose();
        }
    }
}
