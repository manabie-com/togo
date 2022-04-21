package com.manabie.interview.task.controller;

import com.manabie.interview.task.model.Task;
import com.manabie.interview.task.model.User;
import com.manabie.interview.task.response.APIResponse;
import com.manabie.interview.task.service.TaskService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(path = "api/task/")
public class TaskController {

    private final TaskService taskService;

    @Autowired
    public TaskController(TaskService taskService) {
        this.taskService = taskService;
    }

    @GetMapping(path = "/list")
    public ResponseEntity<List<Task>> getTask(){
        List<Task> tasks = taskService.getTasks();
        if(tasks.isEmpty()){
            return new ResponseEntity(HttpStatus.NO_CONTENT);
        }else{
            return new ResponseEntity<>(tasks, HttpStatus.OK);
        }
    }

    @PostMapping(path = "/register")
    public ResponseEntity<APIResponse> registerNewTask(@RequestParam String uid, @RequestParam String createDate, @RequestParam String description){
        Task task = new Task(uid, createDate,description);
        APIResponse response = taskService.registerNewTask(task);
        return new ResponseEntity<APIResponse>(response, response.getStatus());
    }

}
