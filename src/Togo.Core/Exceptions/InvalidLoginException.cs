namespace Togo.Core.Exceptions;

public class InvalidLoginException : Exception
{
    public InvalidLoginException() : base("Username or Password is not valid")
    {
    }
}
