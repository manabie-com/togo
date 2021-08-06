using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace togo.Service.Dto
{
    public class TaskDetailDto
    {
        public string Id { get; set; }
        public string UserId { get; set; }
        public string Content { get; set; }
        public string CreatedDate { get; set; }
    }
}
