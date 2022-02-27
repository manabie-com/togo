using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TODO.Repositories.Data.DBModels
{
    public class Todo
    {
        public int TodoId { get; set; }
        public int UserId { get; set; }
        public int StatusId { get; set; }
        public string TodoName { get; set; }
        public string TodoDescription { get; set; }
        public int Priority { get; set; }
        public DateTime? DateCreated { get; set; }
        public DateTime? DateModified { get; set; }

        // navigation props
        public virtual User User { get; set; }
        public virtual TodoStatus Status { get; set; }
    }
}
