namespace Todo.Storage.Contract;

/* 
 * This interface should be used when a user request to the database is needed.
 */
public class UserRequest
{
    public string FirstName { get; set; } = null!;

    public string LastName { get; set; } = null!;

    public string Todos { get; set; } = null!;

    public int DailyTaskLimit { get; set; }
}