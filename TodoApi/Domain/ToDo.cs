using System.Text.Json.Serialization;

public class ToDo : IEntity
{
    public int Id { get; set; }
    public string Name { get; set; }
    public string Description { get; set; }
    public int UserId { get; set; }
    [JsonIgnore]
    public virtual User User { get; set; }
}