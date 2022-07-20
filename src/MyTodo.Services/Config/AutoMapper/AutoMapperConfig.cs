using System;
using System.Collections.Generic;
using System.Text;
using AutoMapper;

namespace MyTodo.Services.Config.AutoMapper
{
    public class AutoMapperConfig
    {
        public static MapperConfiguration RegisterMappings()
        {
            return new MapperConfiguration(cfg =>
            {
                cfg.AddProfile(new ModelToViewMappingProfile());
                cfg.AddProfile(new ViewToModelMappingProfile());
            });
        }
    }
}
