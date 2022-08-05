using Togo.Core.Entities;

namespace Togo.Core.AppServices.TaskItems.Dtos;

public class TaskItemDto
{
    public int Id { get; set; }

    public string Title { get; set; }

    public TaskItemDto(TaskItem item)
    {
        Id = item.Id;
        Title = item.Title;
    }
}
