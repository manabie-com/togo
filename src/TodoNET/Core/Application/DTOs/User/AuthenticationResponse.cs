namespace Application.DTOs.User
{
    public class AuthenticationResponse
    {
        public string Id { get; set; }
        public string UserName { get; set; }
        public string Email { get; set; }
        public string JWToken { get; set; }
    }
}
