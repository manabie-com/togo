using Manabie.Togo.Core.Commands;
using Manabie.Togo.Data.Entity;

namespace Manabie.Togo.Domain.Commands.UserTask.Create
{
	public class CreatedUserTaskCommand : ICommand<CreateUserTaskResponse>
	{
		public UserTaskEntity UserTaskEntity { get; set; }
	}
}
