using Manabie.Togo.Data.Entity;
using Manabie.Togo.RedisRepository.BaseDatabase;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Manabie.Togo.RedisRepository.Interface
{
    public interface IUserTaskRepositoryRedis : IBaseRepositoryRedis<UserTaskEntity>
    {
        Task<IEnumerable<UserTaskEntity>> GetAllByDay(Guid userId, DateTime taskDate);
    }
}
