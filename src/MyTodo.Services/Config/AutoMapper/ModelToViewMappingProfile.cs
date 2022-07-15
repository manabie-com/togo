using System;
using System.Collections.Generic;
using System.Text;
using AutoMapper;
using MyTodo.Data.Entities;
using MyTodo.Services.ViewModels;

namespace MyTodo.Services.Config.AutoMapper
{
    public class ModelToViewMappingProfile:Profile
    {
        public ModelToViewMappingProfile()
        {
            CreateMap<TodoItem, TodoItemViewModel>();
            CreateMap<Assignment, AssignmentViewModel>();
            CreateMap<AppUser, AppUserViewModel>();
            CreateMap<AppRole, AppRoleViewModel>();
        }
    }
}
