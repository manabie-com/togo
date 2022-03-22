using AutoMapper;
using log4net.Config;
using Microsoft.AspNetCore.Http;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Models;
using Moq;
using Repositories;
using Repositories.Data;
using Repositories.Infrastructure;
using Services.Configs;
using Services.Implementations;
using Services.Interfaces;
using System;
using System.IO;
using System.Reflection;
using Xunit;

namespace IntergationTest
{
    public class BaseService_Test : IDisposable
    {
        protected ServiceProvider ServiceProvider;
        protected DatabaseContext context;
        protected ILoggerFactory loggerFactory;
        protected Mock<ILoggerFactory> _mockLoggerFactory;

        public BaseService_Test()
        {
            var serviceCollection = new ServiceCollection();

            var logRepository = log4net.LogManager.GetRepository(Assembly.GetEntryAssembly());
            XmlConfigurator.Configure(logRepository, new FileInfo("log4net.config"));

            loggerFactory = LoggerFactory.Create(builder =>
            {
                builder.AddLog4Net();
            });

            serviceCollection.AddSingleton(loggerFactory);

            // add AutoMapper
            var mappingConfig = new MapperConfiguration(mc => {
                mc.AddProfile(new MapperProfile());
            });
            IMapper mapper = mappingConfig.CreateMapper();
            serviceCollection.AddSingleton(mapper);

            serviceCollection.AddSingleton<IHttpContextAccessor, HttpContextAccessor>();

            serviceCollection.AddTransient<IUnitOfWork, UnitOfWork>();

            serviceCollection.AddTransient<IRepository<Users>, Repository<Users>>();
            serviceCollection.AddTransient<IUserService, UserService>();

            serviceCollection.AddTransient<IRepository<Tasks>, Repository<Tasks>>();
            serviceCollection.AddTransient<ITaskService, TaskService>();

            var databaseName = Guid.NewGuid().ToString();
            serviceCollection.AddDbContext<DatabaseContext>(options => options.UseInMemoryDatabase(databaseName));

            ServiceProvider = serviceCollection.BuildServiceProvider();

            context = ServiceProvider.GetRequiredService<DatabaseContext>();
            context.Seed();
        }

        public void Dispose()
        {
            if (context != null)
            {
                context.Database.EnsureDeleted();
                context.Dispose();
            }
            GC.SuppressFinalize(this);
        }
    }
}
