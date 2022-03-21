using MediatR;

public class CreateUserCommand : IRequest<User>
{
    public string Name { get; set; }
    public int LimitedTask { get; set; }
}