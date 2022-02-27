using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TODO.Repositories.Models.RequestModels
{
    public class CreateTodoRequest
    {
        public int UserId { get; set; }
        public int StatusId { get; set; }
        public string TodoName { get; set; }
        public string TodoDescription { get; set; }
        public int Priority { get; set; }
        public DateTime? DateCreated { get; set; }
        public DateTime? DateModified { get; set; }
    }
}
