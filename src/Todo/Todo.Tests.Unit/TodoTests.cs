using Xunit;
using FluentAssertions;
using System;

namespace Todo.Tests.Unit;

public class TodoTests
{
    public class Constructor_Should
    {
        [Theory]
        [InlineData(0)]
        [InlineData(-1)]
        public void ThrowException_WhenIdIsOutOfRange(long id)
        {
            Action sut = () => new Todo.Domain.Models.Todo(id, "Sample Name", "Sample Description", DateTime.UtcNow, DateTime.UtcNow, 1);

            sut.Should().Throw<ArgumentOutOfRangeException>();

        }

        [Theory]
        [InlineData("")]
        [InlineData(null)]
        public void ThrowException_WhenNameIsNullOrEmpty(string name)
        {
            Action sut = () => new Todo.Domain.Models.Todo(1, name, "Sample Description", DateTime.UtcNow, DateTime.UtcNow, 1);

            sut.Should().Throw<ArgumentException>();
        }

        [Theory]
        [InlineData("")]
        [InlineData(null)]
        public void ThrowException_WhenDescriptionIsNullOrEmpty(string description)
        {
            Action sut = () => new Todo.Domain.Models.Todo(1, "Sample Name", description, DateTime.UtcNow, DateTime.UtcNow, 1);

            sut.Should().Throw<ArgumentException>();
        }

        [Theory]
        [InlineData(0)]
        [InlineData(-1)]
        public void ThrowException_WhenUserIdIsOutOfRange(long userId)
        {
            Action sut = () => new Todo.Domain.Models.Todo(1, "Sample Name", "Sample Description", DateTime.UtcNow, DateTime.UtcNow, userId);

            sut.Should().Throw<ArgumentOutOfRangeException>();
        }
    }
}