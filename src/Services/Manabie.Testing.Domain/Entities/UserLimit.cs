using Manabie.Testing.Domain.Common;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Testing.Domain.Entities
{
    public class UserLimit : BaseAuditableEntity<int>
    {
        public int TodoLimit { get; set; }
        public int AddedTodo { get; set; }
        public string UserId { get; set; }
        public ICollection<Todo> Todos { get; set; }
    }
}
