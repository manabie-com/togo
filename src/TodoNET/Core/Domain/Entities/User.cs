using Domain.Common;
using System.Collections.Generic;

namespace Domain.Entities
{
    public class User : AuditableBaseEntity
    {
        public string Password { get; set; }
        public int MaxTodo { get; set; }
        public string Email { get; set; }
        public virtual ICollection<Task> Tasks { get; set; }
    }
}
