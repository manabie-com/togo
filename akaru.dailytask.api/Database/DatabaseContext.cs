using System;
using akaru.dailytask.api.Models;
using Microsoft.EntityFrameworkCore;

namespace akaru.dailytask.api.Database
{
    public class DatabaseContext : DbContext
	{
		public DbSet<User> Users { get; set; }
		public DbSet<TodoItem> TodoItems { get; set; }

		protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
		{
			optionsBuilder.UseSqlite("Data Source=database.db");
		}
	}
}

