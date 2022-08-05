namespace Togo.Core.Entities.Common;

public abstract class AuditedEntity : DateEntity
{
    private string _createdBy;

    public string CreatedBy
    {
        get => _createdBy;
        set => _createdBy = value;
    }

    private string _lastModifiedBy;

    public string LastModifiedBy
    {
        get => _lastModifiedBy;
        set => _lastModifiedBy = value;
    }
}
