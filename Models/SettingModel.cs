namespace ManabieTodo.Models
{
    public class SettingModel
    {
        public ConnectionStrings ConnectionStrings { get; set; }
        public string SecretKey { get; set; }
    }

    public class ConnectionStrings
    {
        public string Default { get; set; }
    }
}