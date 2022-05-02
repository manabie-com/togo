using Manabie.Togo.Core.Base;
using Manabie.Togo.Data.Dto;
using Manabie.Togo.Data.Entity;
using Manabie.Togo.Domain.Commands.UserTask.Create;
using Manabie.Togo.Service.Base;
using System;
using System.Threading.Tasks;

namespace Manabie.Togo.Service.Implememt.UserTask
{
	public interface IUserTaskService : IBaseService
	{
		Task<CreateUserTaskResponse> Create(UserTaskDto item);
		Task<ResponseBase> GetAllTaskByDay(GetUserTaskDto getUserTaskDto);
	}
}
