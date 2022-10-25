﻿using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Mvc;

namespace akaru.dailytask.api.Models
{
	[BindProperties]
	public class TodoItem
	{
		public int TodoItemId { get; set; }
		[Required]
		public string Description { get; set; }
		public DateTime TimeStamp { get; set; }
		[Required]
		public int UserId { get; set; }
	}
}

