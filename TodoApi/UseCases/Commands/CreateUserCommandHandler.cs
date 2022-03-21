using FluentValidation;
using MediatR;

public class CreateUserCommandHandler : IRequestHandler<CreateUserCommand, User>
{
    private readonly IUserRepository _userRepository;
    public CreateUserCommandHandler(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    public async Task<User> Handle(CreateUserCommand request, CancellationToken cancellationToken = default(CancellationToken))
    {
        var user = new User
        {
            Name = request.Name,
            LimitedTask = request.LimitedTask
        };

        var currentUser = _userRepository.Add(user);
        await _userRepository.UnitOfWork.SaveChangesAsync();
        return currentUser;
    }
}