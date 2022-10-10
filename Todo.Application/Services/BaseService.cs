using AutoMapper;
using Todo.Application.Interfaces;

namespace Todo.Application.Services
{
    public abstract class BaseService
    {
        protected readonly IMapper Mapper;
        protected readonly IApplicationDbContext DbContext;
        protected BaseService(IMapper mapper, IApplicationDbContext dbContext)
        {
            Mapper = mapper;
            DbContext = dbContext;

        }
    }
}
