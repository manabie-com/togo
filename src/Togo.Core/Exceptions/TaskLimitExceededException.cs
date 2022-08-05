namespace Togo.Core.Exceptions;

public class TaskLimitExceededException : Exception
{
    public TaskLimitExceededException(int taskLimit) : base("Task limit exceeded")
    {
    }
}
