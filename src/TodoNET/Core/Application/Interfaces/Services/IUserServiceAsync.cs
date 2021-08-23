using Application.DTOs.User;
using Application.Wrappers;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Application.Interfaces.Repositories
{
    public interface IUserServiceAsync
    {
        Task<Response<AuthenticationResponse>> AuthenticateAsync(AuthenticationRequest request);
        Task<Response<string>> CreateUserAsync(CreateUserRequest request);
        string Encrypt(string clearText);
        Task<Response<IReadOnlyList<AuthenticationResponse>>> GetUsers();
    }
}
