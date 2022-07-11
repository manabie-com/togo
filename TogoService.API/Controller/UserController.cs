using System;
using System.Threading.Tasks;
using AutoMapper;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using TogoService.API.Dto;
using TogoService.API.Filter;
using TogoService.API.Infrastructure.Helper;
using TogoService.API.Infrastructure.Helper.MessageUtil;
using TogoService.API.Model;
using TogoService.API.Model.Interface;

namespace TogoService.API.Controller
{
    [ApiController]
    [ApiConventionType(typeof(CustomApiConventions))]
    [Route("api/users")]
    [ServiceFilter(typeof(ValidationActionFilter))]
    public class UserController : ControllerBase
    {
        private readonly ILogger _logger;
        private readonly IMapper _mapper;
        private readonly IUnitOfWork _iUnitOfWork;
        private readonly ITodoTaskRepository _iTodoTaskRepository;
        public UserController(ILogger<UserController> logger, IMapper mapper, IUnitOfWork iUnitOfWork, ITodoTaskRepository iTodoTaskRepository)
        {
            _logger = logger;
            _mapper = mapper;
            _iUnitOfWork = iUnitOfWork;
            _iTodoTaskRepository = iTodoTaskRepository;
            _iTodoTaskRepository.SetUnitOfWork(iUnitOfWork);
        }

        [HttpPost("{userId}/tasks")]
        [ProducesResponseType(typeof(CommonResponse<string>), StatusCodes.Status201Created)]
        public async Task<IActionResult> AddTasksForUser(Guid userId, [FromBody] NewTaskRequest requestData)
        {
            CommonResponse<string> response;
            if (requestData.Date.Equals(new DateTime()))
            {
                response = new CommonResponse<string>(StatusCodes.Status400BadRequest, UserControllerErrMsg.MissingTodoDay, null);
                return StatusCode(response.StatusCode, response);
            }

            if (Guid.Empty.Equals(userId))
            {
                response = GenerateCannotFindUserRes(userId);
                return StatusCode(response.StatusCode, response);
            }

            try
            {
                Model.User user = await _iUnitOfWork.GenericRepository<Model.User>().GetById(userId);
                if (user == null)
                {
                    response = GenerateCannotFindUserRes(userId);
                    return StatusCode(response.StatusCode, response);
                }
                else
                {
                    uint addedTasks = ((uint)(await _iTodoTaskRepository.GetAddedTasks(userId, requestData.Date)).Length);
                    uint canAddTasks = user.MaxDailyTasks - addedTasks;
                    if (requestData.Tasks.Length > canAddTasks)
                    {
                        response = new CommonResponse<string>(StatusCodes.Status422UnprocessableEntity, UserControllerErrMsg.ReachoutMaxTaskPerDay, null);
                        return StatusCode(response.StatusCode, response);
                    }

                    TodoTask[] newTasks = new TodoTask[requestData.Tasks.Length];
                    for (int i = 0; i < requestData.Tasks.Length; i++)
                    {
                        newTasks[i] = _mapper.Map<TodoTask>(requestData.Tasks[i]);
                        newTasks[i].UserId = userId;
                        newTasks[i].TodoDay = requestData.Date;
                    }
                    await _iTodoTaskRepository.AddRange(newTasks);
                    await _iUnitOfWork.Save();
                    response = new CommonResponse<string>(StatusCodes.Status201Created,
                        CommonMessages.Ok,
                        CommonMessages.GetSuccessfulAddedItemsMsg(typeof(Model.TodoTask).Name, newTasks.Length));
                }
            }
            catch (Exception ex)
            {
                _logger.LogDebug(ex, ex.Message);
                response = new CommonResponse<string>(StatusCodes.Status500InternalServerError, ex.Message, null);
            }

            return StatusCode(response.StatusCode, response);
        }

        private CommonResponse<string> GenerateCannotFindUserRes(Guid userId)
        {
            return new CommonResponse<string>(StatusCodes.Status404NotFound, CommonMessages.GetCannotFindMsg(typeof(Model.User).Name, userId), null);
        }
    }
}