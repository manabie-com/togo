namespace ManabieTodo.Services
{
    public interface IAuthenticationService
    {
        string? Login(string username, string password);
        string FakeLogin(string name);
        bool Logout(string token);
    }
}