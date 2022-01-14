using System;
using akaru.dailytask.api.Models;
using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace test.akaru.dailytask.api.Test.Models
{
	[TestClass]
	public class TestUser
	{
		private readonly User sut;
		public TestUser()
		{
			sut = new User();
		}

		[TestMethod]
		public void TodoItemShouldBeEmpty()
		{
			Assert.IsNull(sut.Name);
			Assert.AreEqual(0, sut.DailyTaskLimit);
			Assert.AreEqual(0, sut.UserId);
		}

		[TestMethod]
		public void TodoItemShouldSaveValues()
		{
			sut.DailyTaskLimit = 1;
			sut.UserId = 1;
			sut.Name = "test";
			Assert.IsNotNull(sut.Name);
			Assert.AreEqual(1, sut.DailyTaskLimit);
			Assert.AreEqual(1, sut.UserId);
		}
	}
}

