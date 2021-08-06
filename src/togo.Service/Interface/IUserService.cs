using System.Threading.Tasks;

namespace togo.Service.Interface
{
    public interface IUserService
    {
        Task<string> Login(string userId, string password);
    }
}
