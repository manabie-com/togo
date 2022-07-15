using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using MyTodo.Services.Interfaces;
using MyTodo.Services.ViewModels;
using MyTodo.Services.ViewModels.Assignment;
using MyTodo.Services.ViewModels.TodoItem;
using System;
using System.Linq;
using System.Security.Authentication;
using System.Security.Claims;
using System.Threading.Tasks;

namespace MyTodo.BackendApi.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    [Authorize]
    public class TodoItemsController : ControllerBase
    {
        private readonly ITodoItemService _todoItemService;
        private readonly IAssignmentService _assignmentService;

        public TodoItemsController(ITodoItemService todoItemService, IAssignmentService assignmentService)
        {
            this._todoItemService = todoItemService;
            this._assignmentService = assignmentService;
        }
        [HttpGet()]
        public IActionResult GetAll()
        {
            var result = _todoItemService.GetAll();
            return Ok(result);
        }

        [HttpGet("paging")]
        public IActionResult GetAllPaging([FromQuery] TodoItemPagingRequest request)
        {
            var result = _todoItemService.GetAllPaging(request);
            return Ok(result);
        }

        [HttpGet("{id}")]
        public IActionResult GetById(int id)
        {
            var result = _todoItemService.GetById(id);
            if (result == null)
                return BadRequest("Cannot find TodoItem");
            return Ok(result);
        }


        [HttpPost]
        [Consumes("multipart/form-data")]
        public IActionResult Create([FromForm] TodoItemViewModel viewModel)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }
            var result = _todoItemService.Add(viewModel);
            if (result == 0)
                return BadRequest();

            var dotoItem = _todoItemService.GetById(result);

            return CreatedAtAction(nameof(GetById), new { id = result }, dotoItem);
        }

        [HttpPut("{id}")]
        [Consumes("multipart/form-data")]
        public async Task<IActionResult> Update([FromRoute] int id, [FromForm] TodoItemUpdateRequest request)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }
            request.Id = id;
            var affectedResult = _todoItemService.Update(request);
            if (affectedResult == 0)
                return BadRequest();
            return Ok();
        }

        [HttpDelete("{id}")]
        public async Task<IActionResult> Delete(int id)
        {
            var affectedResult = _todoItemService.Remove(id);
            if (affectedResult == 0)
                return BadRequest();
            return Ok();
        }
        #region Assignment
        [HttpPut()]
        [Route("{id}/assign")]
        [Consumes("multipart/form-data")]
        public async Task<IActionResult> Assign([FromRoute] int id, [FromForm] AssignmentCreateRequest request)
        {
            //if (!ModelState.IsValid)
            //{
            //    return BadRequest(ModelState);
            //}
            request.TodoItemId = id;
            request.AssignedUser = GetLoggedUserId();
            var affectedResult = _assignmentService.Add(request);
            if (affectedResult == 0)
                return BadRequest();
            return Ok();
        }

        #endregion
        private Guid GetLoggedUserId()
        {
            if (!User.Identity.IsAuthenticated)
                throw new AuthenticationException();

            var userName = User.Claims.FirstOrDefault(c => c.Type == ClaimTypes.NameIdentifier).Value;

            return Guid.Parse(userName);
        }
    }
}
