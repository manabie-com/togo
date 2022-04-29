using System.Security.Claims;

namespace ManabieTodo.Constants
{
    public static class Tag
    {
        public static string Default { get => "Default"; }
        public static string ConnectionString { get => "ConnectionStrings"; }
    }

    public static class CacheTag
    {
        public static string Token { get => "TOKEN"; }
        public static string CreatedDate { get => "CREATED"; }
    }

    public static class ClaimTag
    {
        public static string AllowedTaskDay { get => "allowed_task_day"; }
        public static string Id { get => ClaimTypes.NameIdentifier; }
        public static string Name { get => ClaimTypes.Name; }
    }
}