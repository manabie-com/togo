using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TODO.Repositories.Data.DBModels
{
    public class TodoStatus
    {
        public int TodoStatusId { get; set; }
        public string StatusName { get; set; }
        public string StatusDescription { get; set; }
    }
}
