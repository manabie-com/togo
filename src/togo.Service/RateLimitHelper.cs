using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace togo.Service
{
    public class RateLimitHelper
    {
        private static readonly ConcurrentDictionary<string, int> _cache = new ConcurrentDictionary<string, int>();

        public static int Peek(string key)
        {
            if (!_cache.ContainsKey(key))
            {
                _cache.TryAdd(key, 0);
                return 0;
            }

            return _cache[key];
        }

        public static int Increase(string key)
        {
            var curr = Peek(key);
            return _cache[key] = curr + 1;
        }
    }
}
