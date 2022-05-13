using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ToDoApp.DTO.Entity
{
    public class User
    {
        public int UserId { get; set; }
        public int DailyLimit { get; set; }
        public List<ToDo> ToDos { get; set; }
    }
}
