using System;
using akaru.dailytask.api.Database;
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
	}
}

