namespace Togo.Infrastructure.Identities.Dtos;

public class CreateUserDto
{
    public string UserName { get; set; }

    public string Password { get; set; }

    public int MaxTasksPerDay { get; set; }
}