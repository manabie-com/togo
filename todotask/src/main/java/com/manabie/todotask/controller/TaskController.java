package com.manabie.todotask.controller;

import com.manabie.todotask.body.AddTaskRequest;
import com.manabie.todotask.body.BaseResponse;
import com.manabie.todotask.entity.Task;
import com.manabie.todotask.exception.ReachLimitException;
import com.manabie.todotask.exception.UserNotFoundException;
import com.manabie.todotask.service.TaskService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.HttpClientErrorException;

@Controller
@RequestMapping("/task")
public class TaskController {
    @Autowired
    TaskService taskService;

    @PutMapping("/add")
    public @ResponseBody ResponseEntity<BaseResponse> addTask(@RequestBody AddTaskRequest request){
        try {
            Task task = taskService.addTask(request);
            return ResponseEntity.ok(new BaseResponse(0, "success").withData(task));
        } catch (ReachLimitException | UserNotFoundException e){
            return ResponseEntity.badRequest().body(new BaseResponse(1, e.getMessage()));
        } catch (Exception e){
            return ResponseEntity.internalServerError().body(new BaseResponse(2, "Internal Server Error"));
        }
    }
}
