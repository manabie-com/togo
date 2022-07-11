using System;
using System.Text.Json.Serialization;

namespace TogoService.API.Dto
{
    public class NewTaskRequest
    {
        public NewTaskRequest() { }

        [JsonPropertyName("Date")]
        public DateTime Date { get; set; }

        [JsonPropertyName("Tasks")]
        public TaskRequest[] Tasks { get; set; }
    }

    public class TaskRequest
    {
        public TaskRequest() { }

        public TaskRequest(string name, string description)
        {
            Name = name;
            Description = description;
        }

        [JsonPropertyName("Name")]
        public string Name { get; set; }

        [JsonPropertyName("Description")]
        public string Description { get; set; }
    }
}