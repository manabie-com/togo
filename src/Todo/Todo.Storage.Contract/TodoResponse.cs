namespace Todo.Storage.Contract;

/* 
 * This class is used when a Todo request to the database is needed.
 */
public class TodoResponse
{
    /*
     * The id of the Todo.
     */
    public long Id { get; set; }

    /*
     * The name of the Todo.
     */
    public string Name { get; set; } = null!;

    /*
     * The description of the Todo.
     */
    public string Description { get; set;} = null!;

    /*
     * The date the Todo was created.
     */
    public DateTime DateCreatedUTC { get; set; }

    /*
     * The date the Todo was modified.
     */
    public DateTime DateModifiedUTC { get; set; }

    /*
     * The id of the user.
     */
    public long UserId { get; set; }
}