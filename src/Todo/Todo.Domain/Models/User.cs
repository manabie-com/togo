namespace Todo.Domain.Models;

public class User
{
    public long Id { get; }

    public string FirstName { get; }

    public string LastName { get; }

    public IEnumerable<Todo> Todos { get; set; } = null!;

    public int DailyTaskLimit { get; }

    public User(long id, string firstName, string lastName, int dailyTaskLimit)
    {
        if (id < 1)
        {
            throw new ArgumentOutOfRangeException("Id must be greater than 1.", nameof(id));
        }

        if (string.IsNullOrEmpty(firstName))
        {
            throw new ArgumentException("First name must not be null or empty.", nameof(firstName));
        }

        if (string.IsNullOrEmpty(lastName))
        {
            throw new ArgumentException("Last name must not be null or empty.", nameof(lastName));
        }

        if (dailyTaskLimit < 1)
        {
            throw new ArgumentOutOfRangeException("Daily task limit should be greater than 0.", nameof(dailyTaskLimit));
        }

        Id = id;
        FirstName = firstName;
        LastName = lastName;
        DailyTaskLimit = dailyTaskLimit;
    }
}