namespace Togo.Infrastructure.Identities;

public class JwtSettings
{
    public string SecurityKey { get; set; }

    public string Issuer { get; set; }

    public string Audience { get; set; }
}
