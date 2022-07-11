using AutoMapper;
using AutoMapper.QueryableExtensions;
using MyTodo.Data.Entities;
using MyTodo.Infrastructure.Interfaces;
using MyTodo.Services.Interfaces;
using MyTodo.Services.ViewModels;
using MyTodo.Services.ViewModels.Common;
using MyTodo.Services.ViewModels.TodoItem;
using MyTodo.Utilities.Exceptions;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace MyTodo.Services.Impl
{
    public class TodoItemService : ITodoItemService
    {
        private readonly IRepository<TodoItem, int> todoItemRepository;
        private readonly IUnitOfWork unitOfWork;
        private readonly IMapper mapper;

        public TodoItemService(IRepository<TodoItem, int> todoItemRepository, IUnitOfWork unitOfWork, IMapper mapper)
        {
            this.todoItemRepository = todoItemRepository;
            this.unitOfWork = unitOfWork;
            this.mapper = mapper;
        }

        public int Add(TodoItemViewModel viewModel)
        {
            var model = mapper.Map<TodoItemViewModel, TodoItem>(viewModel);
            todoItemRepository.Add(model);
            unitOfWork.Commit();
            return model.Id;
        }

        public int Remove(int id)
        {
            todoItemRepository.Remove(id);
            return unitOfWork.Commit();
        }

        public List<TodoItemViewModel> GetAll()
        {
            //return todoItemRepository.FindAll().ProjectTo<TodoItemViewModel>().ToList();
            return mapper.ProjectTo<TodoItemViewModel>(todoItemRepository.FindAll()).ToList();
        }

        public List<TodoItemViewModel> GetAll(string keyword)
        {
            if (!string.IsNullOrEmpty(keyword))
                return mapper.ProjectTo<TodoItemViewModel>(todoItemRepository.FindAll(x => x.Title.Contains(keyword) || x.Description.Contains(keyword))).ToList();
            else
                return mapper.ProjectTo<TodoItemViewModel>(todoItemRepository.FindAll()).ToList();
        }
        public PagedResult<TodoItemViewModel> GetAllPaging(TodoItemPagingRequest request)
        {
            //1. Select join
            var query = todoItemRepository.FindAll();
            //2. filter
            if (!string.IsNullOrEmpty(request.Keyword))
                query = query.Where(x => x.Title.Contains(request.Keyword)||x.Description.Contains(request.Keyword));

            //3. Paging
            int totalRow = query.Count();

            var data = query.Skip((request.PageIndex - 1) * request.PageSize).Take(request.PageSize);

            //4. Mapping
            var dataVM = mapper.ProjectTo<TodoItemViewModel>(data).ToList();
            //4. Select and projection
            var pagedResult = new PagedResult<TodoItemViewModel>()
            {
                TotalRecords = totalRow,
                PageSize = request.PageSize,
                PageIndex = request.PageIndex,
                Items = dataVM
            };
            return pagedResult;
        }

        public TodoItemViewModel GetById(int id)
        {
            return mapper.Map<TodoItem, TodoItemViewModel>(todoItemRepository.FindById(id));
        }


        public int Update(TodoItemUpdateRequest request)
        {
            var getTodoItem = todoItemRepository.FindById(request.Id);
            if (getTodoItem == null) throw new MyTodoException($"Cannot find a product with id: {request.Id}");

            getTodoItem.Title = request.Title;
            getTodoItem.Description = request.Description;
            getTodoItem.Status = request.Status;

            todoItemRepository.Update(getTodoItem);
            return unitOfWork.Commit();
        }

    }
}
