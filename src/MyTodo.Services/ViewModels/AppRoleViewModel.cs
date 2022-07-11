using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Services.ViewModels
{
    public class AppRoleViewModel
    {
        public virtual Guid Id { get; set; }

        public virtual string Name { get; set; }

        public virtual string NormalizedName { get; set; }
    }
}
