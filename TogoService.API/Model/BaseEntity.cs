using Microsoft.EntityFrameworkCore;
using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace TogoService.API.Model
{
    [Index(nameof(IsDeleted))]
    public abstract class BaseEntity
    {
        public BaseEntity()
        {
            IsDeleted = false;
            UpdatedAt = DateTime.UtcNow;
            CreatedAt = DateTime.UtcNow;
        }

        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        [Column("id")]
        public Guid Id { get; set; }

        [Required]
        [Column("isDeleted")]
        public bool IsDeleted { get; set; }

        [Required]
        [Column("updatedAt")]
        public DateTime UpdatedAt { get; set; }

        [Required]
        [Column("createdAt")]
        public DateTime CreatedAt { get; set; }
    }
}
