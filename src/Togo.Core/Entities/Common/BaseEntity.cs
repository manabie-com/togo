namespace Togo.Core.Entities.Common;

public abstract class BaseEntity
{
    private int _id;

    public int Id
    {
        get => _id;
        set => _id = value;
    }
}