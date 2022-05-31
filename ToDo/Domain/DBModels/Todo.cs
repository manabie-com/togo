using ToDo.Api.Domain.Core;

namespace ToDo.Api.Domain.DBModels
{
    public class Todo : Entity
    {
        public Guid UserId { get; set; }
        public int Status { get; set; }
        public string? TodoName { get; set; }
        public string? TodoDescription { get; set; }
        public DateTime? DateCreated { get; set; }
    }
}
