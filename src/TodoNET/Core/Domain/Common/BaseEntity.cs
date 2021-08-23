using System.ComponentModel.DataAnnotations.Schema;

namespace Domain.Common
{
    public abstract class BaseEntity
    {
        [Column("id")]
        public virtual string Id { get; set; }
    }
}
