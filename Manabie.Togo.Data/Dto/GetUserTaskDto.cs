using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Togo.Data.Dto
{
    public class GetUserTaskDto
    {
        public Guid UserId { get; set; }

        public DateTime TaskDate { get; set; }
    }
}
