using AutoMapper;
using Manabie.Togo.Api.Configurations;
using Manabie.Togo.Api.Controllers;
using Manabie.Togo.Core.Base;
using Manabie.Togo.Data.Dto;
using Manabie.Togo.Data.Entity;
using Manabie.Togo.Domain.Commands.UserTask.Create;
using Manabie.Togo.Service.Implememt.UserTask;
using Moq;
using System;
using System.Threading.Tasks;
using Xunit;

namespace Manabie.Togo.Tests
{
	public class UserTasksControllerTest
	{
		private IMapper _mapper;
		Mock<IUserTaskService> _userTaskService;
		UserTasksController _userTasksController;
		public UserTasksControllerTest()
		{
			_userTaskService = new Mock<IUserTaskService>();
			_userTasksController = new UserTasksController(_userTaskService.Object);
			var mapperConfig = new MapperConfiguration(c =>
			{
				c.AddProfile<AutoMapperProfile>();
			});

			_mapper = mapperConfig.CreateMapper();
		}

		[Fact]
		public void GetAllTaskByDayTest()
		{
			GetUserTaskDto getUserTaskDto = new GetUserTaskDto
			{
				UserId = Guid.Parse("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
				TaskDate = new DateTime(2022, 5, 2, 0, 0, 0),
			};
			var response = new ResponseBase();
			var result = _userTaskService.Setup(x => x.GetAllTaskByDay(getUserTaskDto)).Returns(Task.FromResult(response));
			var task = _userTasksController.GetAllTaskByDay(getUserTaskDto);
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
			var result = _userTaskService.Setup(x => x.Create(userTaskDto)).Returns(Task.FromResult(response));
			var task = _userTasksController.Create(userTaskDto);
			Assert.Equal(0, task.Result.Code);
		}
	}
}
