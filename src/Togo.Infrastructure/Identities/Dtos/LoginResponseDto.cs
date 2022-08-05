namespace Togo.Infrastructure.Identities.Dtos;

public class LoginResponseDto
{
    public string UserName { get; set; }
    
    public string AccessToken { get; set; }

    public LoginResponseDto(string userName, string accessToken)
    {
        UserName = userName;
        AccessToken = accessToken;
    }
}
