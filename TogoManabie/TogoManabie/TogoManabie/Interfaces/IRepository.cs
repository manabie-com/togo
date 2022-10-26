using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using TogoManabie.Models;

namespace TogoManabie.Interfaces
{
    public interface IRepository<T> where T : BaseModel
    {
        void Create(T entity);

        Task<T> GetById(int id);
    }
}
