using MediatR;

public class CreateToDoCommand : IRequest<int>
{
    public string Name { get; set; }
    public string Description { get; set; }
    public int UserId { get; set; }
}