package com.manabie.todotaskapplication.controller.task.impl;

import com.manabie.todotaskapplication.common.utils.UserUtils;
import com.manabie.todotaskapplication.controller.task.TaskApi;
import com.manabie.todotaskapplication.data.pojo.task.ResponseListTaskDto;
import com.manabie.todotaskapplication.data.pojo.task.TaskDto;
import com.manabie.todotaskapplication.service.task.TaskService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
@RestController
public class TaskApiController implements TaskApi {
    private static final Logger LOGGER = LoggerFactory.getLogger(TaskApiController.class);

    private TaskService taskService;
    private HttpServletRequest request;

    @Autowired
    public TaskApiController(TaskService taskService, HttpServletRequest request) {
        this.taskService = taskService;
        this.request = request;
    }

    @Override
    @PostMapping("/tasks")
    public ResponseEntity<UUID> createTask(@RequestBody TaskDto task) {
        Optional<String> userId = UserUtils.getUserIdByHttpServletRequest(request);
        if (!userId.isPresent()) {
            return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
        }

        Optional<UUID> uuid = taskService.createTask(task, userId.get());
        if (uuid.isPresent()) {
            return new ResponseEntity<>(uuid.get(), HttpStatus.OK);
        }

        return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @Override
    @GetMapping("/tasks/{id}")
    public ResponseEntity<TaskDto> getTaskById(@PathVariable UUID id) {
        Optional<TaskDto> taskDto = taskService.getTaskById(id);
        if (taskDto.isPresent()) {
            return new ResponseEntity<>(taskDto.get(), HttpStatus.OK);
        }
        return new ResponseEntity<>(HttpStatus.NOT_FOUND);
    }

    @Override
    @PutMapping("/tasks/{id}")
    public ResponseEntity<Void> updateTask(@RequestBody TaskDto taskDto, @PathVariable UUID id) {
        Integer count = taskService.updateTask(taskDto, id).orElse(0);
        if (count <= 0) {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>(HttpStatus.NO_CONTENT);
    }

    @Override
    @DeleteMapping("/tasks/{id}")
    public ResponseEntity<Void> deleteTask(@PathVariable UUID id) {
        taskService.deleteTask(id);
        return new ResponseEntity<>(HttpStatus.NO_CONTENT);
    }

    @Override
    @GetMapping("/tasks")
    public ResponseEntity<ResponseListTaskDto> getTasks() {

        Optional<List<TaskDto>> taskDto = taskService.getTasks();
        if (taskDto.isPresent()) {
            return new ResponseEntity<>(new ResponseListTaskDto(taskDto.get().size(), taskDto.get()), HttpStatus.OK);
        }
        return new ResponseEntity<>(HttpStatus.NOT_FOUND);
    }


}
