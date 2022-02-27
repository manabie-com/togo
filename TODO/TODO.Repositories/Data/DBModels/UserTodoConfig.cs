using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TODO.Repositories.Data.DBModels
{
    public class UserTodoConfig
    {
        public int UserId { get; set; }
        public int DailyTaskLimit { get; set; }

        public virtual User User { get; set; }
    }
}
