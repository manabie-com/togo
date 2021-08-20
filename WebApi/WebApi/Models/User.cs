using System;
using System.Collections.Generic;

namespace WebApi.Models
{
    public partial class User
    {
        public User()
        {
            Tasks = new HashSet<Task>();
        }

        public Guid Id { get; set; }
        public string Password { get; set; }
        public int MaxTodo { get; set; }

        public virtual ICollection<Task> Tasks { get; set; }
    }
}
