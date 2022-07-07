using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using TogoService.API.Infrastructure.Helper.Constant;

namespace TogoService.API.Model
{
    [Table("TodoTask")]
    public class TodoTask : BaseEntity
    {
        public TodoTask() : base() { }

        public TodoTask(string name, string description, DateTime todoDay, Guid userId, User user) : base()
        {
            Name = name;
            Description = description;
            TodoDay = todoDay;
            UserId = userId;
            User = user;
        }

        [Required]
        [Column("name")]
        [MaxLength(EntityConstantsCollection.MaxLengthDataTypeVarchar)]
        public string Name { get; set; }

        [Column("description")]
        [MaxLength(EntityConstantsCollection.MaxLengthDataTypeVarchar)]
        public string Description { get; set; }

        [Required]
        [Column("todoDay")]
        public DateTime TodoDay { get; set; }

        [Column("userId")]
        public Guid UserId { get; set; }
        public virtual User User { get; set; }
    }
}