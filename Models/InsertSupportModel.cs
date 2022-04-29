using Newtonsoft.Json;

namespace ManabieTodo.Models
{
    public class InsertSupportModel
    {
        [JsonProperty("SEQ")]
        public int seq { get; set; }
    }
}