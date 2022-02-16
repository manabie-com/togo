using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging;

namespace ToDoBackend.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class TasksController : ControllerBase
    {
        private readonly ILogger<TasksController> _logger;
        private readonly ToDoContext _dbContext;

        public TasksController(ToDoContext dbContext, ILogger<TasksController> logger)
        {
            _dbContext = dbContext;
            _logger = logger;
        }

        [HttpGet]
        public async Task<IEnumerable<Models.Task>> Get()
        {
            _logger.LogInformation("Reading all tasks from database...");
            return await _dbContext.Tasks.ToListAsync();
        }
    }
}
