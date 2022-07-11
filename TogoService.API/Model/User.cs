using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using TogoService.API.Infrastructure.Helper.Constant;

namespace TogoService.API.Model
{
    [Table("User")]
    public class User : BaseEntity
    {
        public User() : base() { }

        public User(string name, uint maxDailyTasks, IList<TodoTask> tasks) : base()
        {
            Name = name;
            MaxDailyTasks = maxDailyTasks;
            Tasks = tasks;
        }

        [Required]
        [Column("name")]
        [MaxLength(EntityConstantsCollection.MaxLengthDataTypeVarchar)]
        public string Name { get; set; }

        [Column("maxDailyTasks")]
        public uint MaxDailyTasks { get; set; }

        public virtual IList<TodoTask> Tasks { get; set; }
    }
}