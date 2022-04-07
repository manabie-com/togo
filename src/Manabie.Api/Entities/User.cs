using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text.Json.Serialization;

namespace Manabie.Api.Entities;

[Table("User")]
public class User
{
    public User()
    {
    }

    public User(int id, string firstName, string lastName, string username, int maxTodo, string password)
    {
        Id = id;
        FirstName = firstName;
        LastName = lastName;
        Username = username;
        MaxTodo = maxTodo;
        Password = password;
    }

    [Key]
    //[DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int Id { get; set; }
    public string FirstName { get; set; }
    public string LastName { get; set; }
    public string Username { get; set; }
    public int MaxTodo { get; set; } = 5;
    [JsonIgnore]
    public string Password { get; set; }
}
