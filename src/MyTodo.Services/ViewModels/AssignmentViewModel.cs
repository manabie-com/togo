﻿using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Services.ViewModels
{
    public class AssignmentViewModel
    {
        public int Id { get; set; }

        public int TodoItemId { get; set; }

        public Guid UserId { get; set; }

        public virtual TodoItemViewModel TodoItem { get; set; }

        public virtual AppUserViewModel User { get; set; }

        public Guid AssignedUser { get; set; }

        public DateTime AssignedDate { get; set; }
    }
}