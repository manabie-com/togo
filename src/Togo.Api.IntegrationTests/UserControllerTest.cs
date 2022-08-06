using Microsoft.AspNetCore.Mvc.Testing;

namespace Togo.Api.IntegrationTests;

public class UserControllerTest : BaseIntegrationTestClass
{
    public UserControllerTest(WebApplicationFactory<Program> factory) : base(factory)
    {
    }
}
