using System;
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
        public async Task<IActionResult> Get()
        {
            _logger.LogInformation("Reading all tasks from database...");
            return Ok(await _dbContext.Tasks.ToListAsync());
        }

        [HttpPatch]
        [Route("{id}")]
        public async Task<IActionResult> Patch([FromRoute] string id, [FromBody] Models.Task values)
        {
            _logger.LogInformation("Updating an existing task with {Id}...", id);

            var task = await _dbContext.Tasks.FindAsync(id);
            task.Status = values.Status;
            await _dbContext.SaveChangesAsync();

            return Ok(task);
        }

        [HttpPost]
        public async Task<IActionResult> Post([FromBody] Models.Task values)
        {
            _logger.LogInformation("Creating a new task...");
            var entry = _dbContext.Tasks.Add(new Models.Task
            {
                Id = Guid.NewGuid().ToString(),
                Content = values.Content,
                Status = values.Status,
                UserId = values.UserId,
                CreatedDate = DateTimeOffset.Now,
            });
            await _dbContext.SaveChangesAsync();
            return Ok(entry.Entity);
        }

        [HttpDelete]
        public async Task<IActionResult> Delete()
        {
            _logger.LogInformation("Deleting all tasks from database...");
            _dbContext.Tasks.RemoveRange(_dbContext.Tasks);
            await _dbContext.SaveChangesAsync();
            return Ok();
        }

        [HttpDelete]
        [Route("{id}")]
        public async Task<IActionResult> Delete([FromRoute] string id)
        {
            _logger.LogInformation("Deleting an existing task with ID {Id}...", id);
            var task = new Models.Task { Id = id };
            _dbContext.Tasks.Attach(task);
            _dbContext.Tasks.Remove(task);
            await _dbContext.SaveChangesAsync();
            return Ok();
        }
    }
}
