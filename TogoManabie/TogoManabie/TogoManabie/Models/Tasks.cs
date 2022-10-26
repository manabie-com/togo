using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace TogoManabie.Models
{
    public class Tasks : BaseModel
    {
        public User user { get; set; }

        public string content { get; set; }

        public int user_id { get; set; }

        public DateTime created_date { get; set; }
    }
}
