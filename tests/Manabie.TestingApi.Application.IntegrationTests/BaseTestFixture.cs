using NUnit.Framework;
using System.Threading.Tasks;

namespace Manabie.TestingApi.Application.IntegrationTests;

[TestFixture]
public abstract class BaseTestFixture
{
    [SetUp]
    public async Task TestSetUp()
    {
        await Testing.ResetState();
    }
}
