namespace Todo.Contract;

/*
 * This resource is used when the client wants to create a User.
 */
public class CreateUserResource
{
    /*
     * The first name of the User.
     */ 
    public string FirstName { get; set; } = null!;
    
    /*
     * The last name of the User.
     */ 
    public string LastName { get; set;} = null!;

    /*
     * The limit of daily tasks a User can add.
     */
    public int DailyTaskLimit { get; set; }
}