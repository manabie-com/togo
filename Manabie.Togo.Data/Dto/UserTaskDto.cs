using Manabie.Togo.Data.Base;
using System;
namespace Manabie.Togo.Data.Dto
{
    public class UserTaskDto : BaseDto
    {
        public Guid UserId { get; set; }

        public string TaskName { get; set; }

        public string Description { get; set; }

        public DateTime TaskDate { get; set; }
    }
}
