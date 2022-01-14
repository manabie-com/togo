using System;
using System.Collections.Generic;
using System.Linq;
using akaru.dailytask.api.Controllers;
using akaru.dailytask.api.Database;
using akaru.dailytask.api.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace test.akaru.dailytask.api.Test.Controller
{
	[TestClass]
	public class TestUserController
	{
		private readonly UserController sut;
		private readonly DatabaseContext db;

		public TestUserController()
		{
			db = new DatabaseContext();
			sut = new UserController(db);
		}

		[TestInitialize]
		public void CreateDatabase()
        {
			db.Database.EnsureCreated();
        }

		[TestCleanup]
		public void DestroyDatabase()
        {
			db.Database.EnsureDeleted();
        }

		[TestMethod]
		public void ShouldReturnEmptyUsers()
        {
			var users = GetUsersFromResult(sut.Index());
			Assert.AreEqual(0, users.Count());
        }

		[TestMethod]
		public void ShouldGenerateUsers()
        {
			var count = 5;
			var users = GetUsersFromResult(sut.Generate(count));
			Assert.AreEqual(count, users.Count());

			for (int x = 0; x < 10; x++)
            {
				users = GetUsersFromResult(sut.Generate(count));
			}

			users = GetUsersFromResult(sut.Generate(count));
			Assert.AreEqual(60, users.Count());
		}

		[TestMethod]
		public void ShouldClearUsers()
        {
			var count = 5;
			var users = GetUsersFromResult(sut.Generate(count));
			Assert.AreEqual(count, users.Count());

			users = GetUsersFromResult(sut.Clear());
			Assert.AreEqual(0, users.Count());
		}

		[TestMethod]
		public void ShouldAddAndUser()
        {
			var name = "Test Name";
			var user = new User { Name = name };

			sut.Add(user);
			sut.Generate(10);
			var users = GetUsersFromResult(sut.Index());
			Assert.AreEqual(11, users.Count());
        }

		private IEnumerable<User> GetUsersFromResult(IActionResult result)
        {
			var value = ((JsonResult)result).Value;
			if (value is IEnumerable<User>)
            {
				return (IEnumerable<User>)value;
            }
			return new List<User>();
		}
	}
}

