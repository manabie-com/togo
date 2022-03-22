using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Services.ViewModels
{
    public class UserViewModel
    {
        public string ID { get; set; }
        public string Username { get; set; }
        public bool NotFound { get; set; }
    }
}
