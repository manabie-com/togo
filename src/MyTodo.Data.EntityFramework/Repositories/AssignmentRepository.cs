using MyTodo.Data.Entities;
using MyTodo.Data.Interfaces.Repositories;
using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Data.EntityFramework.Repositories
{
    public class AssignmentRepository : EFRepository<Assignment, int>, IAssignmentRepository
    {
        public AssignmentRepository(MyTodoDbContext context) : base(context) { }
    }
}
