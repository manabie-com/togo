using FluentValidation;
using MediatR;

public class CreateToDoCommandHandler : IRequestHandler<CreateToDoCommand, int>
{
    private readonly IUserRepository _userRepository;
    public CreateToDoCommandHandler(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    public async Task<int> Handle(CreateToDoCommand request, CancellationToken cancellationToken = default(CancellationToken))
    {
        var toDo = new ToDo
        {
            Name = request.Name,
            Description = request.Description,
            UserId = request.UserId
        };

        await _userRepository.AddToDoAsync(toDo);
        await _userRepository.UnitOfWork.SaveChangesAsync();
        return toDo.Id;
    }
}