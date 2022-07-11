using AutoMapper;
using TogoService.API.Model;
using TogoService.API.Dto;

namespace TogoService.API.Infrastructure.Mapper
{
    public partial class AutoMapperProfile : Profile
    {
        public AutoMapperProfile()
        {
            CreateMap<TaskRequest, TodoTask>()
               .ForMember(dest => dest.Name, opt => opt.MapFrom(src => src.Name))
               .ForMember(dest => dest.Description, opt => opt.MapFrom(src => src.Description))
               .ForAllOtherMembers(opt => opt.Ignore());
        }
    }
}
