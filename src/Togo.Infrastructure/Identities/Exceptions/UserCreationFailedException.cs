using Microsoft.AspNetCore.Identity;

namespace Togo.Infrastructure.Identities.Exceptions;

public class UserCreationFailedException : Exception
{
    public IdentityResult Result { get; }
    
    public UserCreationFailedException(IdentityResult result) : base("User creation failed")
    {
        Result = result;
    }    
}
