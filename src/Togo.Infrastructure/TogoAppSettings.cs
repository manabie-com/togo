using Togo.Infrastructure.Identities;

namespace Togo.Infrastructure;

public class TogoAppSettings
{
    public JwtSettings JwtBearer { get; set; }

    public ConnectionStrings ConnectionStrings { get; set; }
}

public class ConnectionStrings
{
    public string Default { get; set; }
}
