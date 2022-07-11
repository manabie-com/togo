package com.manabie.todotask.controller;

import com.manabie.todotask.body.AddTaskRequest;
import com.manabie.todotask.body.BaseResponse;
import com.manabie.todotask.entity.Task;
import com.manabie.todotask.exception.ReachLimitException;
import com.manabie.todotask.exception.UserNotFoundException;
import com.manabie.todotask.service.TaskService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.validation.Valid;

@Controller
@RequestMapping("/task")
public class TaskController {
    @Autowired
    TaskService taskService;

    @PutMapping("/add")
    public @ResponseBody ResponseEntity<BaseResponse> addTask(@RequestBody @Valid AddTaskRequest request) {
        try {
            Task task = taskService.addTask(request);
            return ResponseEntity.ok(new BaseResponse(HttpStatus.OK.value(), "success").withData(task));
        } catch (ReachLimitException | UserNotFoundException e) {
            return ResponseEntity.badRequest().body(new BaseResponse(HttpStatus.BAD_REQUEST.value(), e.getMessage()));
        } catch (Exception e) {
            return ResponseEntity.internalServerError().body(new BaseResponse(HttpStatus.INTERNAL_SERVER_ERROR.value(), "Internal Server Error"));
        }
    }
}
