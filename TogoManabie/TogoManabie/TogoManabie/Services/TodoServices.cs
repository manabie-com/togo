using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using TogoManabie.Configuration;
using TogoManabie.Interfaces;
using TogoManabie.Models;
using TogoManabie.Repository;

namespace TogoManabie.Services
{
    public class TodoServices : ITodoServices
    {
        public TodoRepository _todoRepository;

        public UserRepository _userRepository;

        public async Task<Tasks> CreateTodo(Tasks task)
        {
            var user = await _userRepository.GetById(task.user_id);
            var today = DateTime.Now;
            var lstTaskbyUser = await _todoRepository.GetAllByUserId(user.id, today);
            if(lstTaskbyUser.Count > user.maxTodo)
            {
                throw new AppException("You have pass the limit of Task can create in one day");
            }
            _todoRepository.Create(task);
            return task;
        }
    }
}
