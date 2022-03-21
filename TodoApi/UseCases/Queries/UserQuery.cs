using MediatR;

public class UserQuery : IRequest<User>
{
    public int Id { get; set; }
}