using System.Threading.Tasks;

namespace togo.Service.Interface
{
    public interface IUserService
    {
        Task<(bool, string)> Login(string userId, string password);
    }
}
