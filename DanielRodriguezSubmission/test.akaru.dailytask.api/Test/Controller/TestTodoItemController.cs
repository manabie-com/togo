﻿using System;
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
	public class TestTodoItemController
	{
		private readonly TodoItemController sut;
		private readonly DatabaseContext db;
		private readonly User dummyUser;

		public TestTodoItemController()
		{
			db = new DatabaseContext();
			sut = new TodoItemController(db);
			dummyUser = new User { DailyTaskLimit = 5, Name = "Test" };
		}

		// Created and Destroy the Database on each Test
		// This is to ensure test cases are not related and do not depend
		// on the previous test case in order to pass
		[TestInitialize]
		public void CreateDatabase()
		{
			db.Database.EnsureCreated();
			db.Users.Add(dummyUser);
			db.SaveChanges();
		}

		[TestCleanup]
		public void DestroyDatabase()
		{
			db.Database.EnsureDeleted();
		}

		[TestMethod]
		public void ShouldReturnEmptyItemWhenNothingIsAdded()
		{
			var items = GetItemsFromResult(sut.Index());
			Assert.AreEqual(0, items.Count());
		}

		[TestMethod]
		public void ShouldReturnNotFoundErrorWhenUserDoesNotExist()
        {
			var item = new TodoItem { UserId = 234 };
			Assert.IsTrue(sut.Add(item) is NotFoundObjectResult);
        }

		[TestMethod]
		public void ShouldAddTodoItemAndClearItems()
        {
			var item = new TodoItem { UserId = 1, Description = "Test Description" };
			Assert.AreNotEqual(DateTime.Today, item.TimeStamp);
			sut.Add(item);
			var items = GetItemsFromResult(sut.Index());
			Assert.AreEqual(1, items.Count());

			var resultItem = items.First();
			Assert.AreEqual(DateTime.Today, resultItem.TimeStamp);
			Assert.AreEqual(item.Description, resultItem.Description);

			sut.Clear();

			var clearItems = GetItemsFromResult(sut.Index());

			Assert.AreEqual(0, clearItems.Count());
		}

		[TestMethod]
		public void ShouldNotAddWhenDailyLimitReached()
        {
			var user = db.Users.First();
			for (int x = 0; x < user.DailyTaskLimit; x++)
            {
				var item = new TodoItem { UserId = user.UserId, Description = "Test Description" };
				sut.Add(item);
            }
			var newItem = new TodoItem { UserId = user.UserId, Description = "Test Description" };

			Assert.AreEqual(dummyUser.DailyTaskLimit, GetItemsFromResult(sut.Index()).Count());
			Assert.IsTrue(sut.Add(newItem) is BadRequestObjectResult);

		}

		// This is used the parse the JsonResult object we got from the controller
		private IEnumerable<TodoItem> GetItemsFromResult(IActionResult result)
		{
			var value = ((JsonResult)result).Value;
			if (value is IEnumerable<TodoItem>)
			{
				return (IEnumerable<TodoItem>)value;
			}
			return new List<TodoItem>();
		}
	}
}

