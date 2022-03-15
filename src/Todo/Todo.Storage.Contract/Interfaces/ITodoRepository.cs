namespace Todo.Storage.Contract.Interfaces;

/* 
 * This interface should be used when storing Todo data.
*/
public interface ITodoRepository
{
    /*
     * Adds the Todo to the repository.
     */
    Task<TodoResponse> AddAsync(TodoRequest request);
    
    /*
     * Deletes the Todo from the repository based on the supplied Todo id.
     */
    Task<long> DeleteAsync(long id);
}