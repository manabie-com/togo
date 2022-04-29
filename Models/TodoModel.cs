using Newtonsoft.Json;

namespace ManabieTodo.Models
{
    public class TodoModel
    {
        [JsonProperty("ID")]
        public int Id { get; set; } = 0;
        [JsonProperty("DO")]
        public string Do { get; set; } = "";
        [JsonProperty("IS_ACTIVE")]
        public bool IsActive { get; set; } = true;
        [JsonProperty("IS_COMPLETE")]
        public bool IsComplete { get; set; } = false;
        [JsonProperty("CREATED_DATE")]
        public DateTime CreatedDate { get; set; } = DateTime.Now;
        [JsonProperty("UPDATED_DATE")]
        public DateTime? UpdatedDate { get; set; }
        [JsonProperty("ASSIGNEE")]
        public int Assignee { get; set; } = 0;
    }
}