using System;
using akaru.dailytask.api.Database;
using akaru.dailytask.api.Models;
using Microsoft.AspNetCore.Mvc;

namespace akaru.dailytask.api.Controllers
{
	public class UserController : Controller
	{
		private DatabaseContext _db;

		public UserController(DatabaseContext db)
		{
			_db = db;
		}

		public IActionResult Index()
        {
			return Json(_db.Users.ToList());
        }

        public IActionResult Generate(int num)
        {
			var users = Enumerable.Range(0, num).Select(x => GenerateRandomUser());
			_db.AddRange(users);
			_db.SaveChanges();
			return Json(_db.Users.ToList());
        }
		public IActionResult Clear()
        {
			_db.Users.RemoveRange(_db.Users);
			_db.SaveChanges();
			return Redirect("/User");
        }

        private User GenerateRandomUser()
        {
			// Generate User with Random Name and Random DailyTaskLimit
			int min = 5, max = 20;
			return new User
			{
				Name = Faker.Name.FullName(),
				DailyTaskLimit = Random.Shared.Next(min, max)
			};
        }
	}
}

