using System;
using System.Collections.Generic;
using System.Linq;
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
        public async Task<IActionResult> Get([FromQuery] string userId)
        {
            _logger.LogInformation("Reading all tasks from database...");
            return Ok(await _dbContext.Tasks.Where(x => x.UserId == userId).ToListAsync());
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

            if (string.IsNullOrEmpty(values.UserId))
            {
                return Unauthorized();
            }

            var user = await _dbContext.Users.FindAsync(values.UserId);
            if (user is null)
            {
                return Unauthorized();
            }

            await _dbContext.Entry(user).Reference(x => x.Settings).LoadAsync();
            var today = DateTimeOffset.Now.Date;
            if (user.Settings != null)
            {
                var tasksCount = await _dbContext.Tasks
                    .Where(task => task.UserId == user.Id && task.CreatedDate == today)
                    .CountAsync();

                if (tasksCount >= user.Settings.MaxTasksPerDay)
                {
                    _logger.LogInformation("Number of tasks reaches maximum limit per day for user ID {UserId}.", user.Id);
                    return BadRequest("Number of tasks reaches maximum limit per day.");
                }
            }

            var entry = _dbContext.Tasks.Add(new Models.Task
            {
                Id = Guid.NewGuid().ToString(),
                Content = values.Content,
                Status = values.Status,
                UserId = user.Id,
                CreatedDate = DateTimeOffset.Now.Date,
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
