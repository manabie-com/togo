using AutoMapper;
using Microsoft.Extensions.Logging;
using Moq;
using Repositories;
using Repositories.Infrastructure;
using Services.Configs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using UnitTests.Data;

namespace UnitTests
{
    public abstract class BaseService_Test
    {
        protected Mock<DatabaseContext> _mockContext;
        protected Mock<IUnitOfWork> _mockUnitOfWork;
        protected Mock<ILogger> _mockLogger;
        protected Mock<ILoggerFactory> _mockLoggerFactory;
        protected IMapper mapper;

        public BaseService_Test()
        {
            _mockContext = new Mock<DatabaseContext>();
            _mockUnitOfWork = new Mock<IUnitOfWork>();
            _mockLogger = new Mock<ILogger>();
            _mockLoggerFactory = new Mock<ILoggerFactory>();

            //Initializer data
            Initializer initializer = new Initializer();
            initializer.Initialize(_mockContext);

            var mappingConfig = new MapperConfiguration(mc => {
                mc.AddProfile(new MapperProfile());
            });
            mapper = mappingConfig.CreateMapper();

            Setup();
        }

        protected abstract void Setup();
    }
}
