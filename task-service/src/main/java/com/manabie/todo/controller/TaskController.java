package com.manabie.todo.controller;

import com.manabie.todo.model.BaseResponse;
import com.manabie.todo.model.CreateTaskRequest;
import com.manabie.todo.model.TaskModel;
import com.manabie.todo.service.TaskService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/task")
@RequiredArgsConstructor
public class TaskController {
    final TaskService taskService;

    @PostMapping()
    BaseResponse<TaskModel> createTask(@RequestBody CreateTaskRequest request) {
        return BaseResponse.ofSucceeded(taskService.createTask(request));
    }

    @GetMapping
    BaseResponse<List<TaskModel>> findAllTask() {
        return BaseResponse.ofSucceeded(taskService.findAllTask());
    }

    @GetMapping("/owner/{userId}")
    BaseResponse<List<TaskModel>> findAllTaskByOwner(@PathVariable Long userId) {
        return BaseResponse.ofSucceeded(taskService.findAllByOwner(userId));
    }

    @GetMapping("/{id}")
    TaskModel getTaskById(@PathVariable Long id) {
        return BaseResponse.ofSucceeded(taskService.getTaskById(id)).getData();
    }
}