using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Text;
using System.Text.Json.Serialization;

namespace MyTodo.Services.ViewModels.Assignment
{
    public class AssignmentCreateRequest
    {
        [JsonIgnore]
        public int TodoItemId { get; set; }

        public Guid UserId { get; set; }

        public DateTime AssignedDate { get; set; }
        [JsonIgnore]
        public Guid AssignedUser { get; set; }
    }
}
