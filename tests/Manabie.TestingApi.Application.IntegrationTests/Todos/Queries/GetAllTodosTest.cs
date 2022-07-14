using FluentAssertions;
using Manabie.Testing.Application.Todos.Queries.GetAllTodos;
using Manabie.Testing.Domain.Entities;
using NUnit.Framework;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.TestingApi.Application.IntegrationTests.Todos.Queries
{
    public class GetAllTodosTest : BaseTestFixture
    {

        [Test]
        public async Task ShouldReturnAllListsAndItemsByUserId()
        {
            //Arrange
            var data = await Testing.RunAsDefaultUserAsync();
            var userId = data.Item1;
            await Testing.AddRangeAsync(new List<Todo>
            {
                new Todo { Title = "Apples", Note = "Pending", UserId = userId },
                new Todo { Title = "Milk", Note = "Failed", UserId = userId },
                new Todo { Title = "Bread", Note = "Done", UserId = userId },
                new Todo { Title = "Toilet paper", UserId = userId },
                new Todo { Title = "Pasta", UserId = Guid.NewGuid().ToString() },
                new Todo { Title = "Tissues", UserId = Guid.NewGuid().ToString() },
                new Todo { Title = "Tuna", UserId = Guid.NewGuid().ToString() }
            });

            var query = new GetAllTodoQuery()
            {
                UserId = userId
            };

            //Act
            var result = await Testing.SendAsync(query);

            //Assert
            result.Should().HaveCount(4);
        }
    }
}
