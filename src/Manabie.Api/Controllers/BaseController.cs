using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

using Manabie.Api.Auth;
using Manabie.Api.Entities;

using Microsoft.AspNetCore.Mvc;

// For more information on enabling MVC for empty projects, visit https://go.microsoft.com/fwlink/?LinkID=397860

namespace Manabie.Api.Controllers;

[Route("api/[controller]")]
[ApiController]
[Authorize]
public class BaseController : ControllerBase
{
    protected User UserInfo
    {
        get
        {
            return (User)HttpContext.Items["User"];
        }
    }
}

