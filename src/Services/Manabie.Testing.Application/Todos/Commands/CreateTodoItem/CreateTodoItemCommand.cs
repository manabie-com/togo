using Manabie.Testing.Application.Common.Models;
using Manabie.Testing.Application.Interfaces;
using Manabie.Testing.Domain.Entities;
using MediatR;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Testing.Application.Todos.Commands.CreateTodoItem
{
    public class CreateTodoItemCommand : IRequest<Result<int>>
    {
        public string? Title { get; set; }
        public string? Note { get; set; }
        public string? UserId { get; set; }
        public string? Role { get; set; }
    }

    public class CreateTodoItemCommandHandler : IRequestHandler<CreateTodoItemCommand, Result<int>>
    {
        private readonly IManabieDbContext _context;
        private readonly IMediator _mediator;

        public CreateTodoItemCommandHandler(IManabieDbContext context, IMediator mediator)
        {
            _context = context;
            _mediator = mediator;
        }

        public async Task<Result<int>> Handle(CreateTodoItemCommand request, CancellationToken cancellationToken)
        {
            var userLimit = _context.UserLimits.Where(s => s.UserId == request.UserId).FirstOrDefault();

            if(userLimit.LastModified < DateTime.Now.AddDays(-1).Date)
            {
                userLimit.AddedTodo = 0;
            }

            if(userLimit.AddedTodo >= userLimit.TodoLimit)
            {
                return Result<int>.Failure(new string[] { "Exceeded Limit Per Day" }, 0);
            }

            var entity = new Todo
            {
                Title = request.Title,
                Note = request.Note,
                UserId = request.UserId,
            };

            userLimit.AddedTodo++;
            userLimit.LastModified = DateTime.Now;
            _context.Todos.Add(entity);

            await _context.SaveChangesAsync(cancellationToken);

            return Result<int>.Success(entity.Id);
        }
    }
}
