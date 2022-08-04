namespace Togo.Infrastructure.Identities.Dtos;

public class UserDto
{
    public string Id { get; set; }

    public string UserName { get; set; }

    public int MaxTasksPerDay { get; set; }

    public UserDto(AppUser user)
    {
        Id = user.Id;
        UserName = user.UserName;
        MaxTasksPerDay = user.MaxTasksPerDay;
    }
}