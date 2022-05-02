using Manabie.Togo.Data.Base;
using System.Collections.Generic;

namespace Manabie.Togo.RedisRepository.BaseDatabase
{
    public class BaseDtoCompare<T> : IEqualityComparer<T> where T : BaseEntity
    {
        public bool Equals(T x, T y)
        {
            if (x.ID == y.ID)
            {
                return true;
            }
            else
            {
                return false;
            }
        }

        public int GetHashCode(T obj)
        {
            return obj.ID.GetHashCode();
        }
    }
}
