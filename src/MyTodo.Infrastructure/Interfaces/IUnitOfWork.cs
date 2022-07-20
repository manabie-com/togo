using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Infrastructure.Interfaces
{
    public interface IUnitOfWork : IDisposable
    {
        int Commit();
    }
}
