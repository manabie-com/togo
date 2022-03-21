using MediatR;

public class UserQueryHandler : IRequestHandler<UserQuery, User>
{
    private readonly IUserRepository _userRepository;
    public UserQueryHandler(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    public async Task<User> Handle(UserQuery request, CancellationToken cancellationToken = default(CancellationToken))
    {
        var currentUser = await _userRepository.GetByIdAsync(request.Id);
        return currentUser;
    }
}