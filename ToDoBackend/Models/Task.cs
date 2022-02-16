using System;

namespace ToDoBackend.Models
{
    public class Task
    {
        public string Id { get; set; }
        public string Content { get; set; }
        public TaskStatus Status { get; set; }
        public DateTimeOffset CreatedDate { get; set; }
        public string UserId { get; set; }
    }

    public enum TaskStatus
    {
        Active = 1,
        Completed = 2
    }
}
