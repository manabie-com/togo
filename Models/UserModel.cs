using Newtonsoft.Json;

namespace ManabieTodo.Models
{
    public class UserModel
    {
        [JsonProperty("ID")]
        public int Id { get; set; } = 0;
        [JsonProperty("NAME")]
        public string? Name { get; set; }
        [JsonProperty("USERNAME")]
        public string? Username { get; set; }
        [JsonProperty("PASSWORD")]
        public string? Password { get; set; } = null;
        [JsonProperty("ALLOWED_TASK_DAY")]
        public int AllowedTaskDay { get; set; } = (new Random()).Next(1, 10);
    }

    public class LoginModel
    {

    }
}