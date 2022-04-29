using System.Collections.ObjectModel;
using ManabieTodo.Constants;
using ManabieTodo.Models;
using ManabieTodo.Utils;
using Newtonsoft.Json;

namespace ManabieTodo.Services
{
    public class AuthenticationService : IAuthenticationService
    {
        private IFakeCache _fakeCache { get; }
        private JwTokenGenerator _jwTokenGenerator { get; }
        private IDatabaseService _dbService { get; }

        public AuthenticationService(
            JwTokenGenerator jwTokenGenerator,
            IDatabaseService dbService)
        {
            _fakeCache = CacheFactory.FakeCache;
            _jwTokenGenerator = jwTokenGenerator;
            _dbService = dbService;
        }

        public string? Login(string username, string password)
        {

            string cmd = GetUsersByUsername();

            IDictionary<string, object> parameters = new Dictionary<string, object>
            {
                ["USERNAME"] = username
            };

            UserModel user = _dbService.ExecuteObjectReader<UserModel>(cmd, parameters);

            if (user == null)
            {
                return null;
            }

            string hashedPassword = Cryptography.CreateMD5(password);

            if (string.Compare(user.Password, hashedPassword, true) == 0)
            {
                string jwt = _jwTokenGenerator.Create(user);

                _fakeCache.UserToken = new Dictionary<string, object>
                {
                    [CacheTag.Token] = jwt
                };

                return jwt;
            }

            return null;
        }

        public string FakeLogin(string name)
        {
            UserModel user = new UserModel
            {
                Name = name,
                Username = Guid.NewGuid().ToString(),
            };

            string jwt = _jwTokenGenerator.Create(user);

            _fakeCache.UserToken = new Dictionary<string, object>
            {
                [CacheTag.Token] = jwt
            };

            return jwt;
        }

        public bool Logout(string token)
        {
            try
            {
                if (_fakeCache.RemoveUserTokenCache(token) == 0)
                {
                    return false;
                }

                return true;
            }
            catch (Exception ex)
            {
                return false;
            }
        }


        private string GetUsersByUsername()
        {
            return @"
            SELECT
            T.ID,
            T.NAME,
            T.USERNAME,
            T.PASSWORD,
            T.ALLOWED_TASK_DAY
            FROM USERS AS T
            WHERE
            T.USERNAME = $USERNAME
            ";
        }
    }
}