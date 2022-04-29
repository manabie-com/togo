namespace ManabieTodo.Services
{
    public interface IFakeCache
    {
        ICollection<IDictionary<string, object>> UserTokens { get; }

        IDictionary<string, object> UserToken { set; }
        ICollection<IDictionary<string, object>> Tasks { get; }

        IDictionary<string, object> Task { set; }
        int RemoveUserTokenCache(string token);
    }
}