using Microsoft.IdentityModel.Tokens;
using Services.ViewModels;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace WebAPI.Security
{
    public interface IJwtHandler
    {
        string Create(UserViewModel userVM);
        TokenValidationParameters Parameters { get; }
    }
}
