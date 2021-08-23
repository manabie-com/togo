using Application.Interfaces;
using Application.Mappings;
using AutoMapper;
using Domain.Entities;
using Domain.Settings;
using Infrastructure.Persistence.Contexts;
using Infrastructure.Persistence.Repositories;
using Microsoft.AspNetCore.Http;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Options;
using NUnit.Framework;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Security.Claims;
using System.Security.Principal;
using System.Transactions;
using TodoNet.Api.Services;

namespace TestProject.Base
{
    public abstract class TestServiceAsyncBase
    {
        private const string DATABASE_NAME = "TODO_UNIT_TEST";
        protected ApplicationDbContext dbContext;
        protected DbContextOptions<ApplicationDbContext> dbOptions;
        protected IOptions<JWTSettings> jwtOption;
        protected IMapper mapper;
        protected IOptions<AppSettings> appSettings;
        protected IAuthenticatedUserService authenticatedUserService;
        protected IGenericRepositoryAsync<User> genericUserRepositoryAsync;
        protected IGenericRepositoryAsync<Task> genericTaskRepositoryAsync;
        private TransactionScope transaction;
        /// <summary>
        ///     The following method runs before test method
        ///     It can be used for preparing the test data
        ///     Please note that all the data added in RunBeforeTest method need to be cleanup in TestCleanup method.
        /// </summary>
        protected virtual void RunBeforeTest() { }

        /// <summary>
        ///     The following methods is used for clean up the test data which were added to the test database in RunBeforeTest
        ///     method.
        ///     The test data added in [Test] method doesnt need to be clean up here because they will be rollback immediately
        ///     after the [Test] method completed.
        /// </summary>
        protected virtual void TestCleanup() { }
        [SetUp]
        public virtual void Setup()
        {
            // Creating an InMemory DB for unit testing purpose
            dbOptions = new DbContextOptionsBuilder<ApplicationDbContext>()
                          .UseInMemoryDatabase(databaseName: DATABASE_NAME)
                          .Options;
            var jwtSetup = new JWTSettings()
            {
                Key = "C1CF4B7DC4C4175B2218DE4F55CA4",
                Issuer = "IssuerTest",
                Audience = "AudienceTest",
                DurationInMinutes = 60
            };
            jwtOption = Options.Create(jwtSetup);
            appSettings = Options.Create(new AppSettings()
            {
                Secret = "SecrectKey"
            });

            var claims = new List<Claim>()
            {
                new Claim("uid","firstUser")
            };
            var identity = new ClaimsIdentity();
            identity.AddClaims(claims);

            var contextUser = new ClaimsPrincipal(identity);
            HttpContextAccessor httpContextAccessor = new HttpContextAccessor();
            httpContextAccessor.HttpContext = new DefaultHttpContext();
            httpContextAccessor.HttpContext.User = contextUser;
            authenticatedUserService = new AuthenticatedUserService(httpContextAccessor);

            dbContext = new ApplicationDbContext(dbOptions, authenticatedUserService);
            var firstUSer = dbContext.Users.FirstOrDefault(_ => _.Email == "firstUser@gmail.com");
            if (firstUSer == null)
            {
                dbContext.Users.Add(new User
                {
                    Id = "firstUser",
                    Email = "firstUser@gmail.com",
                    Password = "qP0odeFvgjtxYm1S8I8DjvnPROt4OqzPt1iLHswmExA=",
                    MaxTodo = 5,
                });
                dbContext.SaveChanges();
            }

            genericUserRepositoryAsync = new GenericRepositoryAsync<User>(dbContext);
            genericTaskRepositoryAsync = new GenericRepositoryAsync<Task>(dbContext);
            var mockAutoMapper = new MapperConfiguration(mc => mc.AddProfile(new GeneralProfile())).CreateMapper().ConfigurationProvider;
            mapper = new Mapper(mockAutoMapper);

            // Execute custom logic in RunBeforeTest to prepare the necessary data for the test method
            RunBeforeTest();
        }
        [TearDown]
        public void Teardown()
        {
            // Clean up the data added in RunBeforeTest to ensure all tests run in an clean environment.
            TestCleanup();
        }
        protected void RunTestMethodAndRollback(Action testMethod)
        {
            // Create a new transaction scope for the test method
            // This is to ensure that all the data inserted into the database by the test method can be rolled back.
            transaction = new TransactionScope();

            // Execute the test method
            testMethod();

            // Dispose the transaction to rollback all data changes.
            transaction.Dispose();
        }
    }
}
