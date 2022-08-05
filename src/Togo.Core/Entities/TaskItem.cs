using Togo.Core.Entities.Common;

namespace Togo.Core.Entities;

public class TaskItem : AuditedEntity
{
    private string _title;

    public string Title
    {
        get => _title;
        set => _title = value;
    }
    
    public TaskItem()
    {
        // For EF Core only
    }

    private TaskItem(string title)
    {
        Title = title;
    }

    public static TaskItem CreateNew(string title)
    {
        return new TaskItem(title);
    }
}
