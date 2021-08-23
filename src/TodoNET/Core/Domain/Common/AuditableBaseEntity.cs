using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Domain.Common
{
    public abstract class AuditableBaseEntity : BaseEntity
    {
        [Required]
        [Column("created_by")]
        public string CreatedBy { get; set; }
        [Required]
        [Column("created_date")]
        public DateTime CreatedDate { get; set; }
        [Column("last_modified_by")]
        public string LastModifiedBy { get; set; }
        [Column("last_modified_date")]
        public DateTime? LastModifiedDate { get; set; }
    }
}
