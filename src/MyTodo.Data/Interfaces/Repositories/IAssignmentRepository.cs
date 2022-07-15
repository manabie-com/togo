using MyTodo.Data.Entities;
using MyTodo.Infrastructure.Interfaces;
using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Data.Interfaces.Repositories
{
    public interface IAssignmentRepository : IRepository<Assignment, int>
    {
    }
}
