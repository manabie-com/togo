namespace ToDo.Api.Requests
{
    public class CreateUserRequest
    {
        public string? FullName { get; set; }
        public int DailyTaskLimit { get; set; }
    }
}
