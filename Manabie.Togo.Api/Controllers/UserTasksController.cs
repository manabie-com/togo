using Manabie.Togo.Core.Base;
using Manabie.Togo.Data.Dto;
using Manabie.Togo.Domain.Commands.UserTask.Create;
using Manabie.Togo.Service.Implememt.UserTask;
using Microsoft.AspNetCore.Mvc;
using System.Threading.Tasks;

namespace Manabie.Togo.Api.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class UserTasksController : ControllerBase
    {
        private IUserTaskService _userTaskService;
        public UserTasksController(IUserTaskService userTaskService)
        {
            _userTaskService = userTaskService;
        }

        [HttpPost("tasks-by-day")]
        public async Task<ResponseBase> GetAllTaskByDay([FromBody] GetUserTaskDto getUserTaskDto)
        {
            var users = await _userTaskService.GetAllTaskByDay(getUserTaskDto);
            return users;
        }

        [HttpPost("insert")]
        public async Task<CreateUserTaskResponse> Create([FromBody] UserTaskDto userTaskDto)
        {
            var users = await _userTaskService.Create(userTaskDto);

            return users;
        }
    }
}
