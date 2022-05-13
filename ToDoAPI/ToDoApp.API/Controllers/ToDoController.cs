using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using ToDoApp.DTO;
using ToDoApp.DTO.Entity;

namespace ToDoApp.API.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class ToDoController : ControllerBase
    {

        private DatabaseContext _dbContext;

        public ToDoController(DatabaseContext dbContext)
        {
            _dbContext = dbContext;

        }
        // GET: api/<ToDoController>
        [HttpGet]
        public async Task<ActionResult<IEnumerable<ToDo>>> GetToDoList()
        {
            return await _dbContext.Todos.ToListAsync();
        }

        // GET api/<ToDoController>/5
        [HttpGet("{id}")]
        public async Task<ActionResult<ToDo>> GetToDoById(int id)
        {
            var toDo = await _dbContext.Todos.FindAsync(id);
            if(toDo == null)
            {
                return NotFound();
            }
            return toDo;
        }

        // POST api/<ToDoController>
        [HttpPost]
        public async Task<IActionResult> Post([FromBody] ToDo todoItem)
        {
            var limitNumber = _dbContext.Users.Find(todoItem.UserId).DailyLimit;
            var currentAmount = _dbContext.Todos.Where(t => t.Id == todoItem.Id).Count();
            
            if(currentAmount > limitNumber)
            { 
                return BadRequest(); 
            }
            try
            {
                _dbContext.Todos.Add(todoItem);
                await _dbContext.SaveChangesAsync();
                return Ok();
            }
            catch
            {
                return BadRequest();
            }
        }

        // PUT api/<ToDoController>/5
        [HttpPut("{id}")]
        public async Task<IActionResult> Put([FromBody] ToDo newToDoItem)
        {
            var todo = _dbContext.Todos.Find(newToDoItem.Id);
            if(todo == null)
            {
                return NotFound();
            }
            else
            {
                _dbContext.Entry<ToDo>(todo).CurrentValues.SetValues(newToDoItem);
                _dbContext.SaveChanges();

                return Ok();
            }
        }

        // DELETE api/<ToDoController>/5
        [HttpDelete("{id}")]
        public IActionResult Delete(int id)
        {
            var todo = _dbContext.Todos.Find(id);
            if (todo != null)
            {
                _dbContext.Todos.Remove(todo);
                _dbContext.SaveChanges();
                return Ok();

            }
            else
            {
                return NotFound();
            }

        }
    }
}
