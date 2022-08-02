using Microsoft.AspNetCore.Identity;

namespace Togo.Infrastructure.Identities;

public class AppUser : IdentityUser
{
    private int _maxTasksPerDay;

    public int MaxTasksPerDay
    {
        get => _maxTasksPerDay;
        set => _maxTasksPerDay = value;
    }
}