using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using TODO.Repositories.Data;
using TODO.Repositories.Data.DBModels;
using TODO.Repositories.Interfaces;
using TODO.Repositories.Models.RequestModels;

namespace TODO.Repositories.Repositories
{
    public class TodoRepository : ITodoRepository
    {
        private readonly TodoContext _context;

        public TodoRepository(TodoContext context)
        {
            _context = context ?? throw new ArgumentNullException(nameof(context));
        }

        public async Task<Todo> CreateTodo(CreateTodoRequest request)
        {
            await using var transaction = await _context.Database.BeginTransactionAsync();

            try
            {
                var newTodo = new Todo
                {
                    UserId = request.UserId,
                    StatusId = request.StatusId,
                    TodoName = request.TodoName,
                    TodoDescription = request.TodoDescription,
                    Priority = request.Priority,
                    DateCreated = DateTime.UtcNow
                };

                var result = await _context.Todo.AddAsync(newTodo);
                await _context.SaveChangesAsync();
                await transaction.CommitAsync();

                return result.Entity;
            }
            catch (Exception)
            {
                await transaction.RollbackAsync();
                throw;
            }
        }
    }
}
