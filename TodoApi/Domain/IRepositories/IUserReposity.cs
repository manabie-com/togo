public interface IUserRepository : IRepository<ToDo>
{
    public User Add(User user);
    public Task<int> AddToDoAsync(ToDo toDo);
    public Task<bool> UpdateLimitedTaskAsync(int userId, int limitedTask);
    public Task<User> GetByIdAsync(int userId);
}