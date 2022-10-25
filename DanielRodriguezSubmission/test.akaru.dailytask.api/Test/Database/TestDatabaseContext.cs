using System;
using System.Linq;
using akaru.dailytask.api.Database;
using akaru.dailytask.api.Models;
using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace test.akaru.dailytask.api.Test.Database
{
	[TestClass]
	public class TestDatabaseContext
	{
		private readonly DatabaseContext sut;

		public TestDatabaseContext()
		{
			sut = new DatabaseContext();
		}
	}
}

