using Services.Models;
using Services.ViewModels;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Services.Interfaces
{
    public interface IUserService
    {
        UserViewModel Authenticate(UserLogin userLogin);
        bool Register(UserRegister userRegister);
    }
}
