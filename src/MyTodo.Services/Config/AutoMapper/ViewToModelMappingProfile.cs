using AutoMapper;
using System;
using System.Collections.Generic;
using System.Text;
using MyTodo.Data.Entities;
using MyTodo.Services.ViewModels;
using MyTodo.Services.ViewModels.TodoItem;

namespace MyTodo.Services.Config.AutoMapper
{
    public class ViewToModelMappingProfile : Profile
    {
        public ViewToModelMappingProfile()
        {
            //TodoItem
            CreateMap<TodoItemViewModel, TodoItem>()
                .ConstructUsing(x => new TodoItem(x.Title, x.Description, x.Priority, x.Status));
            CreateMap<AssignmentViewModel, Assignment>()
                .ConstructUsing(x => new Assignment(x.TodoItemId, x.UserId, x.AssignedDate));
            CreateMap<AppUserViewModel, AppUser>()
                .ConstructUsing(x => new AppUser(x.Id, x.Email, x.PhoneNumber, x.TaskCount, x.TaskLimit));
            CreateMap<AppRoleViewModel, AppRole>()
                .ConstructUsing(x => new AppRole(x.Id, x.Name, x.NormalizedName));
        }
    }
}
