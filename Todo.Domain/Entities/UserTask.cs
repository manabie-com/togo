using Todo.Domain.Enums;

namespace Todo.Domain.Entities
{
    public class UserTask : BaseEntity
    {
        public string Title { get; set; }
        public TypeTask Type { get; set; }
        public string? Description { get; set; }
        public PriorityTask Priority { get; set; }
    }
}
