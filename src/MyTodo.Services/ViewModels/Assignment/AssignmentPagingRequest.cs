using MyTodo.Services.ViewModels.Common;
using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Services.ViewModels.Assignment
{
    public class AssignmentPagingRequest: PagingRequestBase
    {
        public Guid UserId { get; set; }

    }
}
