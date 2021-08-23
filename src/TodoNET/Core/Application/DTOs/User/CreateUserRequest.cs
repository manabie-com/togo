namespace Application.DTOs.User
{
    public class CreateUserRequest : AuthenticationRequest
    {
        public int MaxTodo { get; set; }
    }
}
