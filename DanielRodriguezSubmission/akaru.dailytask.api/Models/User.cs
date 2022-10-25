﻿using System;
using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Mvc;

namespace akaru.dailytask.api.Models
{
	[BindProperties]
	public class User
	{
		public int UserId { get; set; }
		[Required]
		public string Name { get; set; }
		[Required]
		public int DailyTaskLimit { get; set; }
	}
}

