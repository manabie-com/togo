using System;
using System.Text.Json.Serialization;

namespace akaru.dailytask.api.Models
{
	public class TodoItem
	{
		public int TodoItemId { get; set; }
		public string Description { get; set; }
		public DateTime TimeStamp { get; set; }
		public int UserId { get; set; }
		[JsonIgnore] // Ignore the User, this causes self referencing loops
		public User User { get; set; }	
	}
}

