using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Services.ViewModels.Assignment
{
    public class AssignmentUpdateRequest
    {
        public int Id { get; set; }

        public int TodoItemId { get; set; }

        public string UserName { get; set; }

        public Guid AssignedUser { get; set; }

        public DateTime AssignedDate { get; set; }
    }
}
