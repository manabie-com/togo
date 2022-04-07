using System;
using System.Linq;

using Manabie.Api.Entities;
using Manabie.Api.Insfrastructures;
using Manabie.Api.Models;
using Manabie.Api.Services;
using Manabie.Api.Utilities;

using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Options;

using Moq;

using Xunit;
using Xunit.Sdk;

namespace Manabie.Api.Tests.Services;

public class UserServiceTests
{
    private readonly ManaDbContext _contextMock;
    private readonly IUserService _userServiceMock;
    public UserServiceTests()
    {
        DbContextOptionsBuilder<ManaDbContext> builder = new DbContextOptionsBuilder<ManaDbContext>();
        builder.UseInMemoryDatabase(Guid.NewGuid().ToString());
        _contextMock = new ManaDbContext(builder.Options);

        var appSettings = new Mock<IOptions<AppSettings>>();

        appSettings.Setup(ap => ap.Value).Returns(new AppSettings { Secret = "This is my custom Secret key for authentication." });

        _contextMock.Users.Add(new User
        {
            Id = 1,
            FirstName = "Cuong",
            LastName = "Nguyen Sy Manh",
            MaxTodo = 5,
            Password = "pBNZVHvtfjY3V13BsvZZvl2EdkJPq5xCgJomYUtQdvM=",
            Username = "cuongnsm"
        });
        _contextMock.SaveChanges();

        _userServiceMock = new UserService(_contextMock, appSettings.Object);
    }

    [Fact]
    public void AuthenticateUser()
    {
        AuthenticateRequest request = new AuthenticateRequest
        {
            Username = "cuongnsm",
            Password = "password"
        };

        var authenticateResponse = _userServiceMock.Authenticate(request);

        Assert.Equal("cuongnsm", authenticateResponse.Username);
        Assert.Equal("Cuong", authenticateResponse.FirstName);
        Assert.Equal("Nguyen Sy Manh", authenticateResponse.LastName);
        Assert.Equal(1, authenticateResponse.Id);
    }

    [Fact]
    public void AuthenticateWithWrongPassword()
    {
        AuthenticateRequest request = new AuthenticateRequest
        {
            Username = "cuongnsm",
            Password = "password1"
        };

        var authenticateResponse = _userServiceMock.Authenticate(request);
        Assert.Null(authenticateResponse);
    }


    [Fact]
    public void AuthenticateWithWrongUser()
    {
        AuthenticateRequest request = new AuthenticateRequest
        {
            Username = "cuongnsm1",
            Password = "password"
        };

        var authenticateResponse = _userServiceMock.Authenticate(request);
        Assert.Null(authenticateResponse);
    }

    [Fact]
    public void GetId()
    {
        var id = 1;

        var user = _userServiceMock.GetById(id);
        Assert.NotNull(user);
        Assert.Equal(1, user.Id);
        Assert.Equal("cuongnsm", user.Username);
    }

    [Fact]
    public void GetNonExistUser()
    {
        var id = 2;

        var user = _userServiceMock.GetById(id);
        Assert.Null(user);
    }

    [Fact]
    public void GetAllUsers()
    {
        _contextMock.AddRange(new User(2, "No.2", "Last Name 2", "username2", 4, ""));
        _contextMock.SaveChanges();

        var users = _userServiceMock.GetAll();

        Assert.NotNull(users);
        Assert.Equal(2, users.Count());
    }

}

