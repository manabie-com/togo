using Microsoft.AspNetCore.Mvc;
using ManabieTodo.Services;
using ManabieTodo.Models;
using ManabieTodo.Utils;
using Microsoft.AspNetCore.Authorization;
using System.Security.Claims;
using ManabieTodo.Constants;

namespace ManabieTodo.Controllers
{
    [Authorize]
    public class TodoController : BaseController
    {
        private ITodoService _todoService { get; }
        private IFakeCache _fakeCache { get; }

        public TodoController(ITodoService todoService)
        {
            _todoService = todoService;
            _fakeCache = CacheFactory.FakeCache;
        }

        [HttpGet("{id}")]
        public async Task<IActionResult> Get(int id)
        {
            TodoModel todoModel = await _todoService.GetAsync(id);

            if (todoModel == null)
            {
                return new EmptyResult();
            }

            return new OkObjectResult(todoModel);
        }

        [HttpGet("all")]
        public IActionResult GetAll()
        {
            IAsyncEnumerable<TodoModel> result = _todoService.GetAllAsync();

            if (result == null)
            {
                return new EmptyResult();
            }

            return new OkObjectResult(result);
        }

        [HttpPost]
        [PreventTaskCreate]
        public IActionResult Insert(TodoModel model)
        {
            ClaimsIdentity claimsIdentity = User.Identity as ClaimsIdentity;

            int userId = int.Parse(claimsIdentity.FindFirst(ClaimTag.Id).Value);

            model.Assignee = userId;

            int? newId = _todoService.Insert(model);

            if (newId.HasValue)
            {
                model.Id = newId.Value;
                return new OkObjectResult(model);
            }
            return new ConflictResult();
        }

        [HttpPatch]
        public async Task<IActionResult> Update(TodoModel model)
        {
            bool result = _todoService.Update(model);

            TodoModel todoModel = await _todoService.GetAsync(model.Id);

            if (result)
            {
                return new OkObjectResult(todoModel);
            }

            return new ConflictObjectResult(todoModel);
        }

        [HttpPatch("complete/{id}")]
        public IActionResult ToggleComplete(int id)
        {
            return new OkObjectResult(_todoService.ToggleComplete(id));
        }

        [HttpDelete("{id}")]
        public IActionResult Delete(int id)
        {
            return new OkObjectResult(_todoService.Delete(id));
        }

        [HttpDelete]
        public IActionResult DeleteAll()
        {
            return new OkObjectResult(_todoService.DeleteAll());
        }
    }
}