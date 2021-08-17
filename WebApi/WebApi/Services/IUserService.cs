using System;
using System.Threading.Tasks;
using WebApi.ViewModels;

namespace WebApi.Services
{
    public interface IUserService
    {
        Task<LoginViewModel> Login(Guid userId, string password);
    }
}
