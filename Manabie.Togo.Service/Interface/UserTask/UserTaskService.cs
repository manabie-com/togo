using AutoMapper;
using Manabie.Togo.Core.Base;
using Manabie.Togo.Core.Bus;
using Manabie.Togo.Data.Dto;
using Manabie.Togo.Data.Entity;
using Manabie.Togo.Domain.Commands.UserTask.Create;
using Manabie.Togo.Domain.Commands.UserTask.GetByDay;
using Manabie.Togo.Domain.Events.UserTask.Create;
using Manabie.Togo.RedisRepository.Interface;
using Manabie.Togo.Service.Implememt.UserTask;
using System;
using System.Threading.Tasks;

namespace Manabie.Togo.Service.Interface.UserTask
{
	public class UserTaskService : IUserTaskService
	{
		private IUserTaskRepositoryRedis _userTaskRepositoryRedis;
		private readonly IMediatorHandler _bus;
		private readonly IMapper _mapper;
		public UserTaskService(IUserTaskRepositoryRedis userTaskRepositoryRedis, IMediatorHandler bus, IMapper mapper)
		{
			_userTaskRepositoryRedis = userTaskRepositoryRedis;
			_bus = bus;
			_mapper = mapper;
		}

		public async Task<CreateUserTaskResponse> Create(UserTaskDto item)
		{
			var userTaskEntity = _mapper.Map<UserTaskDto, UserTaskEntity>(item);
			userTaskEntity.ID = Guid.NewGuid();

			// Insert db
			var createCommand = new CreatedUserTaskCommand { UserTaskEntity = userTaskEntity };
			var result = await _bus.SendCommand<CreatedUserTaskCommand, CreateUserTaskResponse>(createCommand);

			// Insert db success => insert into redis
			if (result.Code == 0)
			{
				var eventObj = new CreatedUserTaskEvent { UserTaskEntity = userTaskEntity };
				await _bus.RaiseEvent(eventObj);
			}
			return result;
		}

		public async Task<ResponseBase> GetAllTaskByDay(GetUserTaskDto  getUserTaskDto)
		{
			var command = new GetByDayUserTaskCommand { GetUserTaskDto = getUserTaskDto };
			var result = await _bus.SendCommand<GetByDayUserTaskCommand, GetByDayUserTaskResponse>(command);
			return result;
		}
	}
}
