package com.todo.core.user.application.controller;

import com.todo.core.commons.model.GenericResponse;
import com.todo.core.user.application.dto.UserRegistrationDTO;
import com.todo.core.user.service.TodoUserService;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RequestMapping("/api/v1/users")
@RestController
public class UserController {

    private final TodoUserService todoUserService;

    public UserController(TodoUserService todoUserService) {
        this.todoUserService = todoUserService;
    }

    @PostMapping(value = "/create", consumes = "application/json")
    public GenericResponse<Boolean> createUser(@RequestBody UserRegistrationDTO dto) {
        return this.todoUserService.createUser(dto);
    }

}
