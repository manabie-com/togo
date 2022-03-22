using AutoMapper;
using Models;
using Services.Models;
using Services.ViewModels;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Services.Configs
{
    public class MapperProfile : Profile
    {
        public MapperProfile()
        {
            CreateMap<Users, UserViewModel>();
            CreateMap<UserRegister, Users>();
        }
    }
}
