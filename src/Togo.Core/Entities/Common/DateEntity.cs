namespace Togo.Core.Entities.Common;

public abstract class DateEntity : BaseEntity
{
    private DateTimeOffset _createdAt;

    public DateTimeOffset CreatedAt
    {
        get => _createdAt;
        set => _createdAt = value;
    }

    private DateTimeOffset? _lastModifiedAt;

    public DateTimeOffset? LastModifiedAt
    {
        get => _lastModifiedAt;
        set => _lastModifiedAt = value;
    }
}
