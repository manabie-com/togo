using AutoMapper;
using Microsoft.AspNetCore.Identity;
using MyTodo.Data.Entities;
using MyTodo.Infrastructure.Interfaces;
using MyTodo.Services.Interfaces;
using MyTodo.Services.ViewModels;
using MyTodo.Services.ViewModels.Assignment;
using MyTodo.Services.ViewModels.Common;
using MyTodo.Utilities.Exceptions;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace MyTodo.Services.Impl
{
    public class AssignmentService : IAssignmentService
    {
        private readonly IRepository<Assignment, int> _assignmentRepository;
        private readonly UserManager<AppUser> _userManager;
        private readonly IUnitOfWork _unitOfWork;
        private readonly IMapper _mapper;

        public AssignmentService(IRepository<Assignment, int> assignmentRepository,UserManager<AppUser> userManager, IUnitOfWork unitOfWork, IMapper mapper)
        {
            this._assignmentRepository = assignmentRepository;
            this._userManager = userManager;
            this._unitOfWork = unitOfWork;
            this._mapper = mapper;
        }

        public int Add(AssignmentCreateRequest request)
        {
            //validation
            var user = _userManager.Users.SingleOrDefault(x => x.Id == request.UserId);
            var userLimitTodos = user.TaskLimit;
            var userTodos = _assignmentRepository.FindAll().Count(x => x.UserId == request.UserId && x.AssignedDate == request.AssignedDate);
            if (userTodos >= userLimitTodos)
            {
                throw new MyTodoException("Created enough tasks for the day.");
            }
            var model = _mapper.Map<AssignmentCreateRequest, Assignment>(request);
            _assignmentRepository.Add(model);
            _unitOfWork.Commit();
            return model.Id;
        }

        public int Remove(int id)
        {
            _assignmentRepository.Remove(id);
            return _unitOfWork.Commit();
        }

        public List<AssignmentViewModel> GetAll()
        {
            var data = _assignmentRepository.FindAll();
            var result = _mapper.ProjectTo<AssignmentViewModel>(data).ToList();
            return result;
        }

        public PagedResult<AssignmentViewModel> GetAllPaging(AssignmentPagingRequest request)
        {
            //1. Select join
            var query = _assignmentRepository.FindAll();
            //2. filter
            if (request != null)
                query = query.Where(x => x.UserId == request.UserId);

            //3. Paging
            int totalRow = query.Count();

            var data = query.Skip((request.PageIndex - 1) * request.PageSize).Take(request.PageSize);

            //4. Mapping
            var dataVM = _mapper.ProjectTo<AssignmentViewModel>(data).ToList();
            //4. Select and projection
            var pagedResult = new PagedResult<AssignmentViewModel>()
            {
                TotalRecords = totalRow,
                PageSize = request.PageSize,
                PageIndex = request.PageIndex,
                Items = dataVM
            };
            return pagedResult;
        }

        public AssignmentViewModel GetById(int id)
        {
            return _mapper.Map<Assignment, AssignmentViewModel>(_assignmentRepository.FindById(id));
        }


        public int Update(AssignmentUpdateRequest request)
        {
            var getItem = _assignmentRepository.FindById(request.Id);
            if (getItem == null) throw new MyTodoException($"Cannot find assignment with id: {request.Id}");

            var user = _userManager.FindByNameAsync(request.UserName);

            var countAssignment = _assignmentRepository.FindAll()
                .Where(x => x.UserId == request.AssignedUser && x.AssignedDate == request.AssignedDate).Count();


            //getItem.TodoItemId = request.TodoItemId;
            //getItem.UserId = user.Result.Id;
            getItem.AssignedDate = request.AssignedDate;
            //getItem.AssignedUser = request.AssignedUser;

           
            _assignmentRepository.Update(getItem);
            return _unitOfWork.Commit();
        }
    }
}
