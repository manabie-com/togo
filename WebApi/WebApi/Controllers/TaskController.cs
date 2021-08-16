using Microsoft.AspNetCore.Mvc;
using System.Threading.Tasks;
using WebApi.Requests;

namespace WebApi.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class TaskController : ControllerBase
    {
        public TaskController()
        {
        }

        [HttpPost]
        public async Task<IActionResult> CreateTask([FromBody] CreateTaskRequest createTaskRequest)
        {
            return Ok();
        }
    }
}
