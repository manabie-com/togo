package com.todo.core.user.service;

import com.todo.core.user.model.TodoUser;
import com.todo.core.user.repository.TodoUserRepository;
import org.springframework.stereotype.Service;

@Service
public class TodoUserService implements UserService {

    private final TodoUserRepository todoUserRepository;

    public TodoUserService(TodoUserRepository todoUserRepository) {
        this.todoUserRepository = todoUserRepository;
    }


}
