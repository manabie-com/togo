using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Todo.Domain.Enums;

namespace Todo.Domain.Entities
{
    public class User : BaseEntity
    {
        public string Name { get; set; }
        public GenderType Gender { get; set; }
        public DateTime? Birthday { get; set; }
        public string? PhoneNumber { get; set; }
        public string? Email { get; set; }
        public int LimitTask { get; set; }
        public virtual ICollection<UserTask> UserTasks { get; set; }
    }
}
