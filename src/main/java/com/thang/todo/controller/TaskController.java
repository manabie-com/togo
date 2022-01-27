package com.thang.todo.controller;

import com.thang.todo.config.WebSecurityConfig;
import com.thang.todo.entities.Task;
import com.thang.todo.entities.User;
import com.thang.todo.payload.LoginRequest;
import com.thang.todo.payload.TaskRequest;
import com.thang.todo.repositories.TaskRepository;
import com.thang.todo.repositories.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.validation.Valid;
import java.util.Date;
import java.util.List;

@RestController
@RequestMapping("/api")
public class TaskController {
    @Autowired
    UserRepository userRepository;

    @Autowired
    WebSecurityConfig webSecurityConfig;

    @Autowired
    TaskRepository taskRepository;

    private User getCurrentUser () {
        String username = SecurityContextHolder.getContext().getAuthentication().getPrincipal().toString();
        return userRepository.findByUsername(username).get();
    }

    @GetMapping("/tasks")
    public List<Task> getAllTaskFromUser() {
        Long userId = getCurrentUser().getId();
        List<Task> tasks = taskRepository.findByUserId(userId);
        return tasks;
    }

    @PostMapping("/task")
    public Task createNewTask(@Valid @RequestBody TaskRequest taskRequest) {
//        Long userId = getCurrentUser().getId();
        Task task = new Task(
                taskRequest.getContent(),
                taskRequest.getStatus(),
                new Date(),
                2L
        );
        Task newTask = new Task("test","test",new Date(),2L);
        taskRepository.save(newTask);
        return task;
    }
}