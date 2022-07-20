using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace MyTodo.Services.ViewModels
{
    public class AppUserViewModel
    {
        public virtual Guid Id { get; set; }

        public virtual string Email { get; set; }

        public virtual string PhoneNumber { get; set; }

        public virtual string NormalizedUserName { get; set; }

        public virtual string UserName { get; set; }
        [JsonPropertyName("Number of tasks was assigned")]

        public int TaskCount { get; set; }

        public int TaskLimit { get; set; }

        public virtual ICollection<AssignmentViewModel> Assignments { get; set; }


    }
}
