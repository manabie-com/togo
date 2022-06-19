using SSS.EntityFrameworkCore.Extensions.Entities;

namespace Todo.Domain
{
    public class BaseEntity : IAuditableEntity,IBaseEntity<string>
    {
        public string Id { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedAt { get; set; }
        public string? ModifiedBy { get; set; }
        public DateTime? ModifiedAt { get; set; }
    }
}
