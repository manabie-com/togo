using Manabie.Togo.Core.Events;
using Manabie.Togo.Data.Entity;

namespace Manabie.Togo.Domain.Events.UserTask.Create
{
	public class CreatedUserTaskEvent : Event
	{
		public UserTaskEntity UserTaskEntity { get; set; }
	}
}
