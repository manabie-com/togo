using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Manabie.Api.Entities;

[Table("Task")]
public class Task : EntityBase
{
    public Task()
    {
    }

    public Task(int id, string todo, int userId, DateTime createdAt, User user)
    {
        Id = id;
        Todo = todo;
        UserId = userId;
        CreatedAt = createdAt;
        User = user;
    }

    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int Id { get; set; }
    public string Todo { get; set; }
    public int UserId { get; set; }

    [ForeignKey("UserId")]
    public User User { get; set; }
}
