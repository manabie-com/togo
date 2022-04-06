package com.manabie.todotaskapplication.controller.task;

import com.manabie.todotaskapplication.data.pojo.task.ResponseListTaskDto;
import com.manabie.todotaskapplication.data.pojo.task.TaskDto;
import org.springframework.http.ResponseEntity;

import java.util.List;
import java.util.UUID;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public interface TaskApi {
    ResponseEntity<UUID> createTask(TaskDto task);

    ResponseEntity<TaskDto> getTaskById(UUID id);

    ResponseEntity<Void> updateTask(TaskDto task, UUID id);

    ResponseEntity<Void> deleteTask(UUID id);

    ResponseEntity<ResponseListTaskDto> getTasks();
}
