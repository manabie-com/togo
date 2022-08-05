namespace Togo.Core.Interfaces;

public interface ICurrentUserService
{
    string UserId { get; }
    string SessionId { get; }
    IList<string> RoleIds { get; }
    int MaxTasksPerDay { get; }
}
