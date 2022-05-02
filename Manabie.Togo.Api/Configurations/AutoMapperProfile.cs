using AutoMapper;
using Manabie.Togo.Data.Dto;
using Manabie.Togo.Data.Entity;

namespace Manabie.Togo.Api.Configurations
{
    public class AutoMapperProfile : Profile
    {
        public AutoMapperProfile()
        {
            CreateMap<UserTaskEntity, UserTaskDto>().ReverseMap();
            CreateMap<UserTaskDto, UserTaskEntity>().ReverseMap();
        }
    }
}
