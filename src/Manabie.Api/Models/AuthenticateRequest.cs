using System;
using System.ComponentModel.DataAnnotations;

namespace Manabie.Api.Models;

public class AuthenticateRequest
{
    [Required]
    public string Username { get; set; }

    [Required]
    public string Password { get; set; }
}
