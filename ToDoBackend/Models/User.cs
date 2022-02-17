using System.Collections.Generic;

namespace ToDoBackend.Models
{
    public class User
    {
        public string Id { get; set; }
        public string FirstName { get; set; }
        public string LastName { get; set; }

        public ICollection<Task> Tasks { get; set; }
        public UserSettings Settings { get; set; }
    }
}
