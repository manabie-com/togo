namespace Todo.Storage.Contract;

/* 
 * This interface should be used when a user response from the database is needed.
 */
public class UserResponse
{
    /*
     * The id of the User.
     */
    public long Id { get; set; }

    /*
     * The first name of the User.
     */
    public string FirstName { get; set; } = null!;

    /*
     * The last name of the User.
     */
    public string LastName { get; set; } = null!;

    /*
     * The list of Todos.
     */
    public string Todos { get; set; } = null!;

    /*
     * The user's daily task limit.
     */
    public int DailyTaskLimit { get; set; }
}