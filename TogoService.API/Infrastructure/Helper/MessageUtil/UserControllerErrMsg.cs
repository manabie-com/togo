namespace TogoService.API.Infrastructure.Helper.MessageUtil
{
    public static class UserControllerErrMsg
    {
        public static string MissingTodoDay { get { return "Missing day for to do tasks."; } }
        public static string ReachoutMaxTaskPerDay { get { return "User reaches limitation of number of tasks per day."; } }
    }
}
