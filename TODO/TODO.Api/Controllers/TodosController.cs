using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace TODO.Api.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class TodosController : ControllerBase
    {
        private readonly ILogger<TodosController> _logger;

        public TodosController(ILogger<TodosController> logger)
        {
            _logger = logger;
        }

        [HttpPost]
        public async Task<IActionResult> CreateTodo()
        {
            try
            {
                throw new NotImplementedException();
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }
    }
}
