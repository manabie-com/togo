using Microsoft.AspNetCore.Mvc;
using ToDo.Api.Repositories;
using ToDo.Api.Requests;
using ToDo.Api.Validators;
using TODO.Repositories.Data.DBModels;

namespace ToDo.Api.Controllers
{
    [Route("api/users")]
    [ApiController]
    public class UsersController : ControllerBase
    {
        private readonly IMongoBaseRepository<User> _userRepo;
        private readonly ILogger<UsersController> _logger;
        private readonly CreateUserValidator _validationRules;

        public UsersController(
            ILogger<UsersController> logger,
            IMongoBaseRepository<User> userRepo,
            CreateUserValidator validationRules)
        {
            _userRepo = userRepo;
            _logger = logger;
            _validationRules = validationRules;
        }

        [HttpGet]
        public async Task<IActionResult> GetUsers([FromQuery] Guid userId)
        {
            try
            {
                return Ok(await _userRepo.GetByIdAsync(userId, new CancellationToken()));
            }
            catch (Exception e)
            {
                _logger.LogError(e, e.Message);
                return StatusCode(500, e.Message);
            }
        }
        [HttpPost]
        public async Task<IActionResult> CreateUser([FromBody] CreateUserRequest request)
        {
            var validate = await _validationRules.ValidateAsync(request);
            if (!validate.IsValid)
            {
                throw new ArgumentException("Input is incorect");
            }
            var user = new User
            {
                FullName = request.FullName,
                DailyTaskLimit = request.DailyTaskLimit,
                DateCreated = DateTime.Now
            };
            try
            {
                await _userRepo.AddAsync(user, new CancellationToken());
                return Ok(user);
            }
            catch (Exception e)
            {
                _logger.LogError(e, e.Message);
                return StatusCode(500, e.Message);
            }
        }
    }
}
