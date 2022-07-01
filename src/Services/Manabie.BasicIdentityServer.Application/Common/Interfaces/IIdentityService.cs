using Manabie.BasicIdentityServer.Application.Common.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.BasicIdentityServer.Application.Common.Interfaces
{
    public interface IIdentityService
    {
        public Task<string> GetUserNameAsync(string userId);
        public Task<(Result Result, string UserId)> CreateUserAsync(string userName, string password);
        public Task<bool> IsInRoleAsync(string userId, string role);
        public Task<bool> AuthorizeAsync(string userId, string policyName);
        public Task<Result> DeleteUserAsync(string userId);
    }
}
