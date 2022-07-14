using Manabie.Testing.Application.Interfaces;
using Manabie.Testing.Domain.Entities;
using MediatR;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Testing.Application.Todos.UserLimits.GetUserLimitByUserId
{
    public class GetUserLimitByUserIdQuery : IRequest<UserLimit>
    {
        public string? UserId { get; set; }
    }

    public class GetUserLimitByUserIdQueryHandler : IRequestHandler<GetUserLimitByUserIdQuery, UserLimit>
    {
        private readonly IManabieDbContext _context;

        public GetUserLimitByUserIdQueryHandler(IManabieDbContext context)
        {
            _context = context;
        }

        public Task<UserLimit> Handle(GetUserLimitByUserIdQuery request, CancellationToken cancellationToken)
        {
            return _context.UserLimits.Where(t => t.UserId == request.UserId).FirstOrDefaultAsync();
        }
    }
}
