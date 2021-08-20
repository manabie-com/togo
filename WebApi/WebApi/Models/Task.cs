using System;

namespace WebApi.Models
{
    public partial class Task
    {
        public Guid Id { get; set; }
        public string Content { get; set; }
        public Guid UserId { get; set; }
        public DateTime CreatedDate { get; set; }

        public virtual User User { get; set; }
    }
}
