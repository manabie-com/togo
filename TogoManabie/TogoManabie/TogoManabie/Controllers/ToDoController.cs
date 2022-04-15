using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using TogoManabie.Interfaces;
using TogoManabie.Models;

namespace TogoManabie.Controllers
{
    [Route("api/TodoController")]
    [ApiController]
    public class ToDoController : ControllerBase
    {
        public ITodoServices _todoServices;
        public ToDoController(ITodoServices todoServices)
        {
            _todoServices = todoServices;
        }

        // POST api/values
        [HttpPost]
        public async Task<IActionResult> CreateToDoTask([FromBody] Tasks task)
        {
            var modelCreating = await _todoServices.CreateTodo(task);
            return Created("", modelCreating);
        }

    }
}