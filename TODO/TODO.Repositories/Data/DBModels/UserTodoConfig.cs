using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TODO.Repositories.Data.DBModels
{
    public class UserTodoConfig
    {
        [Key]
        public int UserId { get; set; }
        public int DailyTaskLimit { get; set; }
    }
}
