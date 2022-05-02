using AutoMapper;
using Manabie.Togo.Api.Configurations;
using Manabie.Togo.Core.Base;
using Manabie.Togo.Core.Bus;
using Manabie.Togo.Data.Dto;
using Manabie.Togo.Data.Entity;
using Manabie.Togo.Domain.Commands.UserTask.Create;
using Manabie.Togo.Domain.Commands.UserTask.GetByDay;
using Manabie.Togo.Domain.Events.UserTask.Create;
using Manabie.Togo.RedisRepository.Interface;
using Manabie.Togo.Service.Implememt.UserTask;
using Manabie.Togo.Service.Interface.UserTask;
using Moq;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Xunit;

namespace Manabie.Togo.Tests.Service
{
	public class UserTaskServiceTest
	{
		private IMapper _mapper;
		Mock<IUserTaskRepositoryRedis> _userTaskRepositoryRedis;
		Mock<IMediatorHandler> _bus;
		UserTaskService _userTaskService;
		public UserTaskServiceTest()
		{
			_userTaskRepositoryRedis = new Mock<IUserTaskRepositoryRedis>();
			_bus = new Mock<IMediatorHandler>();
			var mapperConfig = new MapperConfiguration(c =>
			{
				c.AddProfile<AutoMapperProfile>();
			});

			_mapper = mapperConfig.CreateMapper();

			_userTaskService = new UserTaskService(_userTaskRepositoryRedis.Object, _bus.Object, _mapper);
		}

		[Fact]
		public void GetAllTaskByDayTest()
		{
			GetUserTaskDto getUserTaskDto = new GetUserTaskDto
			{
				UserId = Guid.Parse("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
				TaskDate = new DateTime(2022, 5, 2, 0, 0, 0),
			};
			var response = new GetByDayUserTaskResponse();
			_bus.Setup(x => x.SendCommand<GetByDayUserTaskCommand, GetByDayUserTaskResponse>(It.IsAny<GetByDayUserTaskCommand>())).Returns(Task.FromResult(response));
			var task = _userTaskService.GetAllTaskByDay(getUserTaskDto);
			Assert.Equal(0, task.Result.Code);
		}

		[Fact]
		public void CreateTest()
		{
			UserTaskDto userTaskDto = new UserTaskDto
			{
				UserId = Guid.Parse("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
				TaskName = "Task-1",
				Description = "This is task 1",
				TaskDate = new DateTime(2022, 5, 2, 0, 0, 0),
				IsDeleted = false
			};
			var response = new CreateUserTaskResponse();
			var userTaskEntity = _mapper.Map<UserTaskDto, UserTaskEntity>(userTaskDto);
			_bus.Setup(x => x.SendCommand<CreatedUserTaskCommand, CreateUserTaskResponse>(It.IsAny<CreatedUserTaskCommand>())).Returns(Task.FromResult(response));
			_bus.Setup(x => x.RaiseEvent(It.IsAny<CreatedUserTaskEvent>())).Returns(Task.FromResult(response));
			var task = _userTaskService.Create(userTaskDto);
			Assert.Equal(0, task.Result.Code);
		}
	}
}
