namespace Todo.Storage.Contract.Interfaces;

/*
 * This interface should be used when storing User data.
 */
public interface IUserRepository
{
    /*
     * Adds the User to the repository.
     */
    Task<UserResponse> AddAsync(UserRequest resource);

    /*
     * Gets the User from the repository given a user id.
     */
    Task<UserResponse> GetAsync(long id);

    /*
     * Deletes the User from the repository given a user id.
     */
    Task<long> DeleteAsync(long id);
}