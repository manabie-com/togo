using Domain.Common;

namespace Domain.Entities
{
    public class Task : AuditableBaseEntity
    {
        public string Content { get; set; }
        public string UserId { get; set; }
        public virtual User User { get; set; }
    }
}
