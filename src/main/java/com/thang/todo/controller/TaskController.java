package com.thang.todo.controller;

import com.thang.todo.config.WebSecurityConfig;
import com.thang.todo.entities.Task;
import com.thang.todo.entities.User;
import com.thang.todo.payload.TaskRequest;
import com.thang.todo.repositories.TaskRepository;
import com.thang.todo.repositories.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;

import java.util.List;
import java.util.Optional;

@CrossOrigin
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
        String username = SecurityContextHolder.getContext().getAuthentication().getName();
        return userRepository.findByUsername(username).get();
    }

    @GetMapping("/tasks")
    public ResponseEntity<List<Task>> getAllTaskByUserId() {
        Long userId = getCurrentUser().getId();

        return new ResponseEntity<>(taskRepository.findByUserId(userId), HttpStatus.OK);
    }

    @GetMapping("/tasks/{id}")
    public ResponseEntity<Task> getTaskByID(@PathVariable Long id) {
        Long userId = getCurrentUser().getId();
        Optional<Task> taskOptional = taskRepository.findByIdAndUserId(id, userId);
        
        return taskOptional.map(task -> new ResponseEntity<>(task, HttpStatus.OK))
            .orElseGet(() -> new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    @PostMapping("/tasks")
    @Transactional
    public ResponseEntity<?> createTask(@Valid @RequestBody TaskRequest taskRequest) {
        Long userId = getCurrentUser().getId();
        Long maximumTasks = getCurrentUser().getMaximumTasks();
        Long numberOfTasks = taskRepository.countByUserId(userId);
        
        if (numberOfTasks >= maximumTasks) {
            return new ResponseEntity<String>("the number of tasks daily limit is reached", HttpStatus.PRECONDITION_FAILED);
        }

        Task task = new Task(
            taskRequest.getContent(),
            taskRequest.getStatus(),
            userId
        );

        return new ResponseEntity<Task>(taskRepository.save(task), HttpStatus.OK);
    }

    @PutMapping("/tasks/{id}")
    public ResponseEntity<Task> updateTask(@PathVariable Long id, @RequestBody TaskRequest taskRequest) {
        Task task = taskRepository.getById(id);
        task.setContent(taskRequest.getContent());
        task.setStatus(taskRequest.getStatus());

        return new ResponseEntity<Task>(taskRepository.save(task), HttpStatus.OK);
    }

    @DeleteMapping("/tasks/{id}")
    public void deleteTask(@PathVariable Long id) {
        taskRepository.deleteById(id);
    }
}