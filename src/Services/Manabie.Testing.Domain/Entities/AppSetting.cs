using Manabie.Testing.Domain.Common;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Testing.Domain.Entities
{
    public class AppSetting : BaseAuditableEntity<int>
    {
        public string Key { get; set; } 
        public string Value { get; set; }
    }
}
