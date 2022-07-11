using MyTodo.Services.ViewModels;
using MyTodo.Services.ViewModels.Assignment;
using MyTodo.Services.ViewModels.Common;
using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Services.Interfaces
{
    public interface IAssignmentService
    {
        List<AssignmentViewModel> GetAll();

        PagedResult<AssignmentViewModel> GetAllPaging(AssignmentPagingRequest request);

        AssignmentViewModel GetById(int id);

        int Add(AssignmentViewModel request);

        int Update(AssignmentUpdateRequest request);

        int Remove(int id);
    }
}
