using ToDo.Api.Domain.Core;

namespace TODO.Repositories.Data.DBModels
{
    public class User : Entity
    {
        public string? FullName { get; set; }
        public int DailyTaskLimit { get; set; }
        public DateTime DateCreated { get; set; }

    }
}
