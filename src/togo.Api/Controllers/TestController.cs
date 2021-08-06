using Microsoft.AspNetCore.Mvc;

namespace togo.Api.Controllers
{
    [Route("[controller]")]
    public class TestController : ControllerBase
    {
        [HttpGet("Ping")]
        public string Ping() => "Pong";
    }
}
