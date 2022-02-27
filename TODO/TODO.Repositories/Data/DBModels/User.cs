using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TODO.Repositories.Data.DBModels
{
    public class User
    {
        public int UserId { get; set; }
        public string LastName { get; set; }
        public string FirstName { get; set; } 
        public string MiddleName { get; set; }

        public virtual ICollection<Todo> Todos { get; set; }
    }
}
