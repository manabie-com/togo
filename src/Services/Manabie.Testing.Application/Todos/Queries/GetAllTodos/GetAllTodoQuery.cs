using Manabie.Testing.Application.Interfaces;
using Manabie.Testing.Domain.Entities;
using MediatR;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Testing.Application.Todos.Queries.GetAllTodos
{
    public class GetAllTodoQuery : IRequest<List<Todo>>
    {
        public string? UserId { get; set; }
    }

    public class GetAllTodoQueryQueryHandler : IRequestHandler<GetAllTodoQuery, List<Todo>>
    {
        private readonly IManabieDbContext _context;

        public GetAllTodoQueryQueryHandler(IManabieDbContext context)
        {
            _context = context;
        }

        public Task<List<Todo>> Handle(GetAllTodoQuery request, CancellationToken cancellationToken)
        {
            return _context.Todos.Where(t => t.UserId == request.UserId).ToListAsync();
        }
    }
}
