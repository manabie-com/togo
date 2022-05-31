using Microsoft.AspNetCore.Mvc;
using MongoDB.Driver;
using ToDo.Api.Domain.DBModels;
using ToDo.Api.Repositories;
using ToDo.Api.Requests;
using ToDo.Api.Validators;
using TODO.Repositories.Data.DBModels;

namespace ToDo.Api.Controllers
{
    [Route("api/todos")]
    [ApiController]
    public class TodosController : ControllerBase
    {
        private readonly ILogger<TodosController> _logger;
        private readonly IMongoBaseRepository<Todo> _todoRepo;
        private readonly IMongoBaseRepository<User> _userRepo;
        private readonly CreateToDoValidator _validationRules;

        public TodosController(
            ILogger<TodosController> logger,
            IMongoBaseRepository<Todo> todoRepo,
            IMongoBaseRepository<User> userRepo,
            CreateToDoValidator validationRules)
        {
            _logger = logger;
            _todoRepo = todoRepo;
            _userRepo = userRepo;
            _validationRules = validationRules;
        }

        [HttpPost]
        public async Task<IActionResult> CreateTodo([FromBody] CreateTodoRequest request)
        {
            var validate = await _validationRules.ValidateAsync(request);
            if (!validate.IsValid)
            {
                return BadRequest("Input is incorect");
            }
            try
            {
                var user = await _userRepo.GetByIdAsync(request.UserId, new CancellationToken());

                if (user == null)
                    return NotFound("User does not exist.");

                var taskLimit = user.DailyTaskLimit;
                var filter = Builders<Todo>.Filter.And(Builders<Todo>.Filter.Eq(x => x.UserId , request.UserId), 
                    Builders<Todo>.Filter.Gte(x => x.DateCreated , DateTime.Now.Date));
                
                var tasks =await _todoRepo.FindAllWithFilter(filter);
                var taskCount = tasks.Count();

                if (taskCount >= taskLimit)
                    return BadRequest("Unable to create TODO: User has exceeded daily TODOs.");

                Todo todonew = new Todo
                {
                    UserId = request.UserId,
                    Status = request.Status,
                    TodoName = request.TodoName,
                    TodoDescription = request.TodoDescription,
                    DateCreated = DateTime.Now
                };

                await _todoRepo.AddAsync( todonew, new CancellationToken());

                return Ok(todonew);
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }
        [HttpGet]
        public async Task<IActionResult> GetTodos([FromQuery] Guid ToDoId)
        {
            try
            {
                return Ok(await _todoRepo.GetByIdAsync(ToDoId, new CancellationToken()));
            }
            catch (Exception e)
            {
                _logger.LogError(e, e.Message);
                return StatusCode(500, e.Message);
            }
        }
    }
}
