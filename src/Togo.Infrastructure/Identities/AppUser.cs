using Microsoft.AspNetCore.Identity;

namespace Togo.Infrastructure.Identities;

public class AppUser : IdentityUser
{
    private int _maxTasksPerDay;

    public AppUser()
    {
        // For EF only
    }

    private AppUser(string userName, int maxTasksPerDay) : base(userName)
    {
        _maxTasksPerDay = maxTasksPerDay;
    }

    public int MaxTasksPerDay
    {
        get => _maxTasksPerDay;
        set => _maxTasksPerDay = value;
    }

    public static AppUser Create(string username, int maxTasksPerDay)
    {
        return new AppUser(username, maxTasksPerDay);
    }
}