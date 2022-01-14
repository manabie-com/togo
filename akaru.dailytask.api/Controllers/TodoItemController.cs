using System;
using akaru.dailytask.api.Database;
using akaru.dailytask.api.Models;
using Microsoft.AspNetCore.Mvc;

namespace akaru.dailytask.api.Controllers
{
	public class TodoItemController : Controller
	{
		private DatabaseContext _db;

		public TodoItemController(DatabaseContext db)
		{
			_db = db;
		}

		public IActionResult Index()
        {
			return Json(_db.TodoItems.ToList());
        }

		[HttpPost]
		public IActionResult Add([FromBody]TodoItem todoItem)
        {
			var userId = todoItem.UserId;
			var user = _db.Users.Find(userId);
			if (user is null)
            {
				return NotFound($"UserId : {userId} not found");
            }

			var currentTasksToday = _db.TodoItems.Count(i => i.UserId == userId && i.TimeStamp == DateTime.Today);

			if (currentTasksToday >= user.DailyTaskLimit)
			{
				return BadRequest($"UserId : {user.UserId} has reached Daily Task Limit of {user.DailyTaskLimit}");
			}
			todoItem.TimeStamp = DateTime.Today;
			_db.TodoItems.Add(todoItem);
			_db.SaveChanges();

			return Json(todoItem);
        }

		public IActionResult Clear()
        {
			_db.TodoItems.RemoveRange(_db.TodoItems);
			_db.SaveChanges();
			return Redirect("/TodoItem"); 
        }
	}
}

