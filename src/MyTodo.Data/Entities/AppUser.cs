using Microsoft.AspNetCore.Identity;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;

namespace MyTodo.Data.Entities
{
    [Table("AppUsers")]
    public class AppUser : IdentityUser<Guid>
    {
        public AppUser()
        {

        }
        public AppUser(Guid userId, string email, string phoneNumber, int taskCount, int taskLimit)
        {
            this.Id = userId;
            this.Email = email;
            this.PhoneNumber = phoneNumber;
            this.TaskCount = taskCount;
            this.TaskLimit = taskLimit;
        }
        public int TaskCount { get; set; }
        public int TaskLimit { get; set; }
        public virtual ICollection<Assignment> Assignments { get; set; }


    }
}
