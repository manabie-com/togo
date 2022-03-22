using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Services.Models
{
    public class UserRegister
    {
        public string Username { get; set; }
        public string Password { get; set; }
        public int TaskPerDay { get; set; }
    }
}
