package com.todo.core.todo.application.controller;

import com.todo.core.commons.model.GenericResponse;
import com.todo.core.commons.utils.JwtUtils;
import com.todo.core.todo.model.Todo;
import com.todo.core.todo.service.TodoService;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.data.web.PageableDefault;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;

@RequestMapping("/api/v1/todos")
@RestController
public class TodoController {

    private final TodoService todoService;
    private final JwtUtils jwtUtils;

    public TodoController(TodoService todoService,
                          JwtUtils jwtUtils) {
        this.todoService = todoService;
        this.jwtUtils = jwtUtils;
    }

    @GetMapping("/list")
    public GenericResponse<Page<Todo>> getAllByUser(HttpServletRequest request,
                                                    @PageableDefault(size = 10, page = 0, sort = "dateCreated",
                                 direction = Sort.Direction.DESC) Pageable pageable) {
        final String jwt = jwtUtils.getJwtFromRequest(request);
        final Long todoUserId = jwtUtils.getUserIdFromJwt(jwt);
        return todoService.retrieveAllTodoByUser(todoUserId, pageable);
    }
}
