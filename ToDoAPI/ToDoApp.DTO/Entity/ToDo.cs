using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ToDoApp.DTO.Entity
{
    public class ToDo
    {
        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public int Id { get; set; }
        [Required(ErrorMessage ="Title is Required")]
        public string Title { get; set; }
        [Required(ErrorMessage = "Detail is Required")]
        public string Detail { get; set; }
        public DateTime CreatedDate { get; set; } = DateTime.Now;
        [Required(ErrorMessage = "User Id is Required")]
        public int UserId { get; set; }
        [ForeignKey("UserId")]
        public virtual User User { get; set; }
    }
}
