namespace Todo.Contract;

/*
 * This resource is used when the client wants to create a Todo.
 */
public class CreateTodoResource
{
    /*
     * To create a Todo for a specified user, a UserId must be supplied.
     */
    public long UserId { get; set; }

    /*
     * The name of the Todo.
     */
    public string Name { get; set; } = null!;

    /*
     * The description of the Todo.
     */
    public string Description { get; set; } = null!;
}