using System;
using akaru.dailytask.api.Models;
using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace test.akaru.dailytask.api.Test.Models
{
	[TestClass]
	public class TestTodoItem
	{
		private readonly TodoItem sut;
		public TestTodoItem()
		{
			sut = new TodoItem();
		}

		[TestMethod]
		public void TodoItemShouldBeEmpty()
        {
			Assert.AreEqual(DateTime.MinValue, sut.TimeStamp);
			Assert.IsNull(sut.Description);
			Assert.AreEqual(0, sut.TodoItemId);
			Assert.AreEqual(0, sut.UserId);
        }

		[TestMethod]
		public void TodoItemShouldSaveValues()
		{
			sut.TodoItemId = 1;
			sut.UserId = 1;
			sut.TimeStamp = DateTime.Today;
			sut.Description = "test";
			Assert.AreEqual(DateTime.Today, sut.TimeStamp);
			Assert.IsNotNull(sut.Description);
			Assert.AreEqual(1, sut.TodoItemId);
			Assert.AreEqual(1, sut.UserId);
		}
	}
}

