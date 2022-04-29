using System.Collections.ObjectModel;
using ManabieTodo.Constants;
using ManabieTodo.Utils;

namespace ManabieTodo.Services
{
    public class FakeCache : IFakeCache
    {
        public ICollection<IDictionary<string, object>> UserTokens { get; }

        public IDictionary<string, object> UserToken
        {
            set
            {
                IDictionary<string, object> userToken = value;
                userToken.Add("CREATE_DATE", DateTime.Now);

                UserTokens.Add(userToken);
            }
        }

        public ICollection<IDictionary<string, object>> Tasks { get; }

        public IDictionary<string, object> Task
        {
            set
            {
                IDictionary<string, object> task = value;
                task.Add("CREATE_DATE", DateTime.Now.ToString("YYYY-MM-dd"));

                Tasks.Add(task);
            }
        }

        public FakeCache()
        {
            UserTokens = new Collection<IDictionary<string, object>>();
            Tasks = new Collection<IDictionary<string, object>>();
        }

        public int RemoveUserTokenCache(string token)
        {
            var userToken = UserTokens.FirstOrDefault(ut => ut.GetString(CacheTag.Token) == token);

            if (userToken == null)
            {
                return 0;
            }

            if (UserTokens.Remove(userToken))
            {
                return 1;
            }

            return 0;
        }
    }
}