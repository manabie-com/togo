package com.manabie.todo.api;

import com.manabie.todo.api.request.AddTaskRequest;
import com.manabie.todo.api.response.AddTaskResponse;
import com.manabie.todo.domain.Task;
import com.manabie.todo.domain.User;
import com.manabie.todo.service.TaskService;
import com.manabie.todo.utils.AppResponse;
import lombok.AllArgsConstructor;
import org.springframework.http.MediaType;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

@RestController
@AllArgsConstructor
@RequestMapping(value = "/task", consumes = MediaType.APPLICATION_JSON_VALUE)
public class TaskController {

  private final TaskService taskService;

  @PostMapping()
  public AppResponse<AddTaskResponse> add(@RequestBody @Validated AddTaskRequest request) {
    var user = (User) SecurityContextHolder.getContext().getAuthentication().getPrincipal();

    var task = taskService.add(user,
        Task.builder()
            .title(request.getTitle())
            .description(request.getDescription())
            .build()
    );

    return AppResponse.ok(
        AddTaskResponse.builder()
            .id(task.getId())
            .title(task.getTitle())
            .description(task.getDescription())
            .build()
    );
  }

  @DeleteMapping()
  public AppResponse<Object> delete(@RequestParam Long id) {
    taskService.delete(
        Task.builder()
            .id(id)
            .build()
    );

    return AppResponse.ok(null);
  }
}
