﻿using MyTodo.Data.Enums;
using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Services.ViewModels
{
    public class TodoItemViewModel
    {
        public int Id { get; set; }
        public string Title { get; set; }
        public string Description { get; set; }
        public int Priority { get; set; }
        public TodoItemStatus Status {get;set;}
        public virtual ICollection<AssignmentViewModel> Assignments { get; set; }
    }
}