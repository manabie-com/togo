using Microsoft.EntityFrameworkCore;

public class UserRepository : IUserRepository
{
    private readonly ToDoDbContext _context;
    private readonly ILogger<UserRepository> _logger;

    public IUnitOfWork UnitOfWork
    {
        get
        {
            return _context;
        }
    }

    public UserRepository(ToDoDbContext toDoDbContext, ILogger<UserRepository> logger)
    {
        _context = toDoDbContext ?? throw new ArgumentNullException(nameof(toDoDbContext));
        _logger = logger;
    }

    public User Add(User user)
    {
        return _context.Users
                .Add(user)
                .Entity;
    }

    public async Task<int> AddToDoAsync(ToDo toDo)
    {
        var currentUser = await _context.Users.SingleOrDefaultAsync(x => x.Id == toDo.UserId);
        if (currentUser == null || currentUser.TotalTask == currentUser.LimitedTask)
        {
            throw new Exception("Can not update, try it again");
        }

        currentUser.TotalTask = currentUser.TotalTask + 1;
        _context.Entry(currentUser).State = EntityState.Modified;
        _context.ToDos.Add(toDo);
        return toDo.Id;
    }

    public async Task<bool> UpdateLimitedTaskAsync(int userId, int limitedTask)
    {
        var currentUser = await _context.Users.SingleOrDefaultAsync(x => x.Id == userId);
        if (currentUser != null)
        {
            try
            {
                currentUser.LimitedTask = limitedTask;
                _context.Entry(currentUser).State = EntityState.Modified;
                return true;
            }
            catch (Exception ex)
            {
                _logger.LogError(ex?.Message);
                return false;
            }
        }

        return false;
    }

    public async Task<User> GetByIdAsync(int userId)
    {
        var currentUser = await _context.Users.Where(x => x.Id == userId).Include(x => x.ToDos).FirstOrDefaultAsync();
        return currentUser;
    }
}