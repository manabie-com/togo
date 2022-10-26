using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace TogoManabie.Models
{
    public class User : BaseModel
    {
        public IList<Tasks> tasks { get; set; }

        public string passWord { get; set; }

        public int maxTodo { get; set; }
    }
}
