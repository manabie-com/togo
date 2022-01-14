using System;
using System.Text.Json.Serialization;

namespace akaru.dailytask.api.Models
{
	public class User
	{
		public int UserId { get; set; }
		public string Name { get; set; }
		public int DailyTaskLimit { get; set; }
		[JsonIgnore] // Ignore the User, this causes self referencing loops
		public List<TodoItem> TodoItems { get; set; } = new List<TodoItem>();
	}
}

