using Microsoft.AspNetCore.Identity;
using System;
using System.ComponentModel.DataAnnotations.Schema;

namespace MyTodo.Data.Entities
{
    [Table("AppRoles")]
    public class AppRole : IdentityRole<Guid>
    {
        public AppRole()
        {

        }
        public AppRole(Guid roleId, string roleName, string normalizedName)
        {
            this.Id = roleId;
            this.Name = roleName;
            this.NormalizedName = normalizedName;
        }
    }
}
