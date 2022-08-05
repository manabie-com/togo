namespace Togo.Core.Exceptions;

public class TaskLimitExceededException : Exception
{
    public int TaskLimit { get; }
    
    public TaskLimitExceededException(int taskLimit) : base("Task limit exceeded")
    {
        TaskLimit = taskLimit;
    }
}
