

using ManabieTodo.Services;

namespace ManabieTodo.Utils
{
    public static class CacheFactory
    {
        private static IFakeCache _fakeCache { get; set; }

        public static IFakeCache FakeCache
        {
            get
            {
                if (_fakeCache == null)
                {
                    _fakeCache = new FakeCache();
                }

                return _fakeCache;
            }
        }
    }
}