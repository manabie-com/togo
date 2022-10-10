using System;
using System.ComponentModel.DataAnnotations;
using AutoMapper;
using Todo.Application.Mapping;
using Todo.Domain.Entities;
using Todo.Domain.Enums;

namespace Todo.Application.Dtos
{
    public class CreateEditTaskDto : IMapTo<UserTask>
    {
        //this field will be gotten from claims if user has authorization and authentication
        [Required]
        public string UserId { get; set; }
        [Required]
        public string Title { get; set; }
        public string Description { get; set; }
        public PriorityTask Priority { get; set; }
        public TypeTask Type { get; set; }
        public void Mapping(Profile profile)
        {
            profile.CreateMap<CreateEditTaskDto, UserTask>()
                    .ForMember(x => x.CreatedBy, opts => opts.MapFrom(c => c.UserId));
        }
    }
}
