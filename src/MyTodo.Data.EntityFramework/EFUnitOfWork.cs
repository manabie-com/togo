using MyTodo.Infrastructure.Interfaces;
using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Data.EntityFramework
{
    public class EFUnitOfWork : IUnitOfWork
    {
        private readonly MyTodoDbContext _context;
        public EFUnitOfWork(MyTodoDbContext context)
        {
            _context = context;
        }
        public int Commit()
        {
            return _context.SaveChanges();
        }

        public void Dispose()
        {
            _context.Dispose();
        }
    }
}
