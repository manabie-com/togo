namespace ToDoBackend.Models
{
    public class UserSettings
    {
        public string UserId { get; set; }
        public int MaxTasksPerDay { get; set; }

        public User User { get; set; }
    }
}
