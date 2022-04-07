using System;
using System.Linq;

using Manabie.Api.Entities;
using Manabie.Api.Insfrastructures;
using Manabie.Api.Services;
using Manabie.Api.Utilities;

using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Options;

using Moq;

using Xunit;

namespace Manabie.Api.Tests.Services;

public class TaskServiceTests
{
    private readonly ManaDbContext _contextMock;
    private readonly ITaskService _taskService;
    public TaskServiceTests()
    {
        DbContextOptionsBuilder<ManaDbContext> builder = new DbContextOptionsBuilder<ManaDbContext>();
        builder.UseInMemoryDatabase(Guid.NewGuid().ToString());
        _contextMock = new ManaDbContext(builder.Options);

        var appSettings = new Mock<IOptions<AppSettings>>();

        appSettings.Setup(ap => ap.Value).Returns(new AppSettings { Secret = "This is my custom Secret key for authentication." });

        _contextMock.Users.AddRange(new User
        {
            Id = 1,
            FirstName = "Cuong",
            LastName = "Nguyen Sy Manh",
            MaxTodo = 5,
            Password = "pBNZVHvtfjY3V13BsvZZvl2EdkJPq5xCgJomYUtQdvM=",
            Username = "cuongnsm"
        },
        new User
        {
            Id = 2,
            FirstName = "Cuong",
            LastName = "Nguyen Sy Manh",
            MaxTodo = 2,
            Password = "pBNZVHvtfjY3V13BsvZZvl2EdkJPq5xCgJomYUtQdvM=",
            Username = "test1"
        },
        new User
        {
            Id = 3,
            FirstName = "Cuong",
            LastName = "Nguyen Sy Manh",
            MaxTodo = 4,
            Password = "pBNZVHvtfjY3V13BsvZZvl2EdkJPq5xCgJomYUtQdvM=",
            Username = "test2"
        },
        new User
        {
            Id = 4,
            FirstName = "Cuong",
            LastName = "Nguyen Sy Manh",
            MaxTodo = 9,
            Password = "pBNZVHvtfjY3V13BsvZZvl2EdkJPq5xCgJomYUtQdvM=",
            Username = "test3"
        });

        _contextMock.AddRange(
        new Task
        {
            Todo = "Task 1",
            UserId = 1
        },
        new Task
        {
            Todo = "Task 2",
            UserId = 1
        },
        new Task
        {
            Todo = "Task 3",
            UserId = 1
        },
        new Task
        {
            Todo = "Task 4",
            UserId = 1
        },
        new Task
        {
            Todo = "Task 1",
            UserId = 2
        },
        new Task
        {
            Todo = "Task 2",
            UserId = 2
        },
        new Task
        {
            Todo = "Task 3",
            UserId = 4
        },
        new Task
        {
            Todo = "Task 4",
            UserId = 3
        }
        );


        _contextMock.SaveChanges();

        _taskService = new TaskService(_contextMock);
    }

    [Fact]
    public void GetUserTasks()
    {
        var userName = "cuongnsm";
        var result = _taskService.GetTasks(userName);
        Assert.NotNull(result);
        Assert.NotEmpty(result);
        Assert.Equal(4, result.Count());
    }

    [Fact]
    public void GetTaskWithWrongUsers()
    {
        var userName = "cuongnsm9";
        var result = _taskService.GetTasks(userName);
        Assert.Null(result);
    }

    [Fact]
    public void AddTaskInNormalCase()
    {
        var userName = "cuongnsm";
        var result = _taskService.AddTask(userName, new Models.TaskViewModel { Todo = "Task 5" });
        Assert.True(result.Success);
        Assert.True(string.IsNullOrEmpty(result.Message));

        var col = _taskService.GetTasks(userName);
        Assert.NotNull(col);
        Assert.NotEmpty(col);
        Assert.Equal(5, col.Count());
    }

    [Fact]
    public void AddTaskInAbnormalCase()
    {
        var userName = "test1";
        var result = _taskService.AddTask(userName, new Models.TaskViewModel { Todo = "Task 3" });
        Assert.False(result.Success);
        Assert.Equal("Your list is maximized.", result.Message);

        var col = _taskService.GetTasks(userName);
        Assert.NotNull(col);
        Assert.NotEmpty(col);
        Assert.Equal(2, col.Count());
    }

    [Fact]
    public void AddTaskWithWrongUser()
    {
        var userName = "test5";
        var result = _taskService.AddTask(userName, new Models.TaskViewModel { Todo = "Task 3" });
        Assert.False(result.Success);
        Assert.Equal("This user is not exist in the system.", result.Message);
    }

    [Fact]
    public void AddTaskWithEmptyOrNullTodo()
    {
        var userName = "test3";
        var result = _taskService.AddTask(userName, new Models.TaskViewModel { Todo = String.Empty });
        Assert.False(result.Success);
        Assert.Equal("Todo can't be empty.", result.Message);

        var col = _taskService.GetTasks(userName);
        Assert.NotNull(col);
        Assert.NotEmpty(col);
        Assert.Equal(1, col.Count());
    }
}

