using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using MyTodo.Services.Interfaces;
using MyTodo.Services.ViewModels;
using MyTodo.Services.ViewModels.TodoItem;
using System.Threading.Tasks;

namespace MyTodo.BackendApi.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class TodoItemsController : ControllerBase
    {
        private readonly ITodoItemService todoItemService;

        public TodoItemsController(
            ITodoItemService todoItemService)
        {
            this.todoItemService = todoItemService;
        }
        [HttpGet()]
        //[Authorize]
        public IActionResult GetAll()
        {
            var result = todoItemService.GetAll();
            return Ok(result);
        }

        [HttpGet("paging")]
        public IActionResult GetAllPaging([FromQuery] TodoItemPagingRequest request)
        {
            var result = todoItemService.GetAllPaging(request);
            return Ok(result);
        }

        [HttpGet("{id}")]
        public IActionResult GetById(int id)
        {
            var result = todoItemService.GetById(id);
            if (result == null)
                return BadRequest("Cannot find TodoItem");
            return Ok(result);
        }


        [HttpPost]
        [Consumes("multipart/form-data")]
        [Authorize]
        public IActionResult Create([FromForm] TodoItemViewModel viewModel)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }
            var result = todoItemService.Add(viewModel);
            if (result == 0)
                return BadRequest();

            var dotoItem = todoItemService.GetById(result);

            return CreatedAtAction(nameof(GetById), new { id = result }, dotoItem);
        }

        [HttpPut("{id}")]
        [Consumes("multipart/form-data")]
        [Authorize]
        public async Task<IActionResult> Update([FromRoute] int id, [FromForm] TodoItemUpdateRequest request)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }
            request.Id = id;
            var affectedResult = todoItemService.Update(request);
            if (affectedResult == 0)
                return BadRequest();
            return Ok();
        }

        [HttpDelete("{id}")]
        [Authorize]
        public async Task<IActionResult> Delete(int id)
        {
            var affectedResult = todoItemService.Remove(id);
            if (affectedResult == 0)
                return BadRequest();
            return Ok();
        }
    }
}
