namespace togo.Service.Interface
{
    public interface IJwtGenerator
    {
        string CreateToken(string userId);
    }
}
