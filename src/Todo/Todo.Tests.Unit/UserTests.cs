using Xunit;
using FluentAssertions;
using System;

namespace Todo.Tests.Unit;

public class UserTests
{
    public class Constructor_Should
    {
        [Theory]
        [InlineData(0)]
        [InlineData(-1)]
        public void ThrowException_WhenIdIsOutOfRange(long id)
        {
            Action sut = () => new Todo.Domain.Models.User(id, "FirstName", "LastName", 10);

            sut.Should().Throw<ArgumentOutOfRangeException>();
        }

        [Theory]
        [InlineData("")]
        [InlineData(null)]
        public void ThrowException_WhenFirstNameIsNullOrEmpty(string firstName)
        {
            Action sut = () => new Todo.Domain.Models.User(1, firstName, "LastName", 10);

            sut.Should().Throw<ArgumentException>();
        }

        [Theory]
        [InlineData("")]
        [InlineData(null)]
        public void ThrowException_WhenLastNameIsNullOrEmpty(string lastName)
        {
            Action sut = () => new Todo.Domain.Models.User(1, "FirstName", lastName, 10);

            sut.Should().Throw<ArgumentException>();
        }

        [Theory]
        [InlineData(0)]
        [InlineData(-1)]
        public void ThrowException_WhenDailyTaskLimitIsOutOfRange(int dailyTaskLimit)
        {
            Action sut = () => new Todo.Domain.Models.User(1, "FirstName", "LastName", dailyTaskLimit);

            sut.Should().Throw<ArgumentOutOfRangeException>();
        }
    }
}