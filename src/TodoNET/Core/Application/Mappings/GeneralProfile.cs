using Application.DTOs.Task;
using Application.DTOs.User;
using AutoMapper;
using Domain.Entities;

namespace Application.Mappings
{
    public class GeneralProfile : Profile
    {
        public GeneralProfile()
        {
            CreateMap<CreateTaskRequest, Task>();
            CreateMap<CreateTaskRequest, Task>().ReverseMap();
            CreateMap<TaskResponse, Task>();
            CreateMap<TaskResponse, Task>().ReverseMap();

            CreateMap<User, AuthenticationResponse>();
            CreateMap<User, AuthenticationResponse>().ReverseMap();
            CreateMap<User, CreateUserRequest>();
            CreateMap<User, CreateUserRequest>().ReverseMap();
        }
    }
}
