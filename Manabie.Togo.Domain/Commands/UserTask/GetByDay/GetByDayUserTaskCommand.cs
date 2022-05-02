using Manabie.Togo.Core.Commands;
using Manabie.Togo.Data.Dto;

namespace Manabie.Togo.Domain.Commands.UserTask.GetByDay
{
	public class GetByDayUserTaskCommand : ICommand<GetByDayUserTaskResponse>
	{
		public GetUserTaskDto GetUserTaskDto { get; set; }
	}
}
