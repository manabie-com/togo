using Manabie.Testing.Domain.Common;
using System.ComponentModel.DataAnnotations.Schema;
using System.Diagnostics.CodeAnalysis;

namespace Manabie.Testing.Domain.Entities
{
    public class Todo : BaseAuditableEntity<int>
    {
        public string? Title { get; set; }
        public string? Note { get; set; }
        public string UserId { get; set; }
        [ForeignKey("UserLimitId")]
        public int UserLimitId { get; set;}
        public virtual UserLimit UserLimit { get; set; }

    }
}
