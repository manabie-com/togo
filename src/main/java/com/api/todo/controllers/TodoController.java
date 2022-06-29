package com.api.todo.controllers;

import javax.validation.Valid;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.WebDataBinder;
import org.springframework.web.bind.annotation.InitBinder;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.api.todo.entities.Task;
import com.api.todo.request.RequestTaskEntity;
import com.api.todo.services.TodoService;
import com.api.todo.validators.TaskValidator;

@RestController
@RequestMapping("/api/todo")
public class TodoController {
    @Autowired
    private TodoService todoService;

    @Autowired
    private TaskValidator taskValidator;

    @InitBinder(value = "requestTaskEntity")
    void initTaskValidator(WebDataBinder binder) {
        binder.setValidator(taskValidator);
    }

    @PostMapping("/tasks")
    public ResponseEntity<Task> createTask(@RequestBody @Valid RequestTaskEntity requestTaskEntity) {
        Task task = buildEntityFromRequest(requestTaskEntity);

        // create task
        Task savedTask = todoService.createTask(task);
        return new ResponseEntity<Task>(savedTask, HttpStatus.CREATED);
    }

    private Task buildEntityFromRequest(@Valid RequestTaskEntity requestTaskEntity) {
        Task task = new Task(requestTaskEntity);
        return task;
    }
}
