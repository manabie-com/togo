namespace Todo.Domain.Models;

public class Todo
{
    public long Id { get; }

    public string Name { get; }

    public string Description { get; }

    public DateTime DateCreatedUTC { get; }

    public DateTime DateModifiedUTC { get; }

    public long UserId { get; }

    public Todo(long id, string name, string description, DateTime dateCreatedUTC, DateTime dateModifiedUTC, long userId)
    {
        if (id < 1)
        {
            throw new ArgumentOutOfRangeException("Id must be greater than 1.", nameof(id));
        }

        if (userId < 1)
        {
            throw new ArgumentOutOfRangeException("UserId must be greater than 1.", nameof(UserId));
        }

        if (string.IsNullOrEmpty(name))
        {
            throw new ArgumentException("Name must not be null or empty.", nameof(name));
        }

        if (string.IsNullOrEmpty(description))
        {
            throw new ArgumentException("Description must not be null or empty.", nameof(description));
        }

        Id = id;
        Name = name;
        Description = description;
        DateCreatedUTC = dateCreatedUTC;
        DateModifiedUTC = dateCreatedUTC;
        UserId = userId;
    }
}