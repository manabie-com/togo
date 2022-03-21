public class User : IEntity
{
    public int Id { get; set; }
    public string Name { get; set; }
    public int LimitedTask { get; set; }
    public int TotalTask { get; set;}
    public virtual ICollection<ToDo> ToDos { get; set; }
}