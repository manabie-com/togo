using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Models
{
    public class Tasks
    {
        public string ID { get; set; }
        public string Content { get; set; }
        public DateTime CreateAt { get; set; }

        public string UserID { get; set; }
        [ForeignKey("UserID")]
        public virtual Users User { get; set; }
    }
}
