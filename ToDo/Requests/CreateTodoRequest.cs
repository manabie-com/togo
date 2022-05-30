namespace ToDo.Api.Requests
{
    public class CreateTodoRequest
    {
        public Guid UserId { get; set; }
        public Guid ToDoID { get; set; }
        public int Status { get; set; }
        public string? TodoName { get; set; }
        public string? TodoDescription { get; set; }
    }
}
