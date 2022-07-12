using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace MyTodo.BackendApi.Models
{
    public class UpdateUserViewModel
    {
        public Guid UserId { get; set; }
        public int TaskLimit { get; set; }
    }
}
