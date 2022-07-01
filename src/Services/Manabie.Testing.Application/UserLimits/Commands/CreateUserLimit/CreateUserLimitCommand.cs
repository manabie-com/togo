using Manabie.Testing.Application.Interfaces;
using Manabie.Testing.Domain.Entities;
using MediatR;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Testing.Application.UserLimits.Commands.CreateUserLimit
{
    public class CreateUserLimitCommand : IRequest<int>
    {
        public string UserId { get; set; }
        public string Role { get; set; }
    }
    public class CreateUserLimitCommandHandler : IRequestHandler<CreateUserLimitCommand, int>
    {
        private readonly IManabieDbContext _context;

        public CreateUserLimitCommandHandler(IManabieDbContext context)
        {
            _context = context;
        }

        public async Task<int> Handle(CreateUserLimitCommand request, CancellationToken cancellationToken)
        {
            var entity = new UserLimit() { UserId = request.UserId, LastModified = DateTime.Now, Created = DateTime.Now };

            switch (request.Role)
            {
                case "Administration":
                    entity.TodoLimit = 5;
                    entity.AddedTodo = 0;
                    break;
                case "User":
                    entity.TodoLimit = 3;
                    entity.AddedTodo = 0;
                    break;
                default:
                    entity.TodoLimit = 3;
                    entity.AddedTodo = 0;
                    break;
            }
            await _context.UserLimits.AddAsync(entity);

            await _context.SaveChangesAsync();

            return entity.Id;
        }
    }
}
