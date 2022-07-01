using FluentAssertions;
using Manabie.Testing.Application.Todos.Commands.CreateTodoItem;
using Manabie.Testing.Application.UserLimits.Commands.CreateUserLimit;
using Manabie.Testing.Domain.Entities;
using NUnit.Framework;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.TestingApi.Application.IntegrationTests.Todos.Commands
{
    public class CreateTodosTest : BaseTestFixture
    {
        [Test]
        public async Task ShouldCreateTodoItem()
        {
            var data = await Testing.RunAsDefaultUserAsync();

            var userId = data.Item1;
            var role = data.Item2;

            var userLimitId = await Testing.SendAsync(new CreateUserLimitCommand()
            {
                UserId = userId,
                Role = role,
            });

            var result = await Testing.SendAsync(new CreateTodoItemCommand
            {
                Title = "New Todo",
                Note = "Pending",
                UserId = userId,
                Role = role,
            });

            var item = await Testing.FindAsync<Todo>(result.Data);

            item.Should().NotBeNull();
            item.UserId.Should().Be(userId);
        }

        [TestCase("User", 3)]
        [TestCase("Adminitrator", 5)]
        public async Task ShouldCreateNotExceededLimit(string role, int limit)
        {
            //Arrange
            var data = await Testing.RunAsDefaultUserAsync();

            var userId = data.Item1;
            var userLimitId = await Testing.SendAsync(new CreateUserLimitCommand()
            {
                UserId = userId,
                Role = role,
            });

            for (int i = 0; i < limit; i++)
            {
                await Testing.SendAsync(new CreateTodoItemCommand
                {
                    Title = "New Todo " + i,
                    Note = "Pending",
                    UserId = userId,
                    Role = role,
                });
            }

            //Act
            var result = await Testing.SendAsync(new CreateTodoItemCommand
            {
                Title = "Exceeded Todo",
                Note = "Pending",
                UserId = userId,
                Role = role,
            });

            var item = await Testing.FindAsync<Todo>(result.Data);

            //Assert
            result.Succeeded.Should().Be(false);
            result.Errors.Should().HaveCountGreaterThan(0);
            item.Should().BeNull();
        }
    }
}
