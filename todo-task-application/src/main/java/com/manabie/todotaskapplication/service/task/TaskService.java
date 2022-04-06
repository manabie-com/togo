package com.manabie.todotaskapplication.service.task;

import com.manabie.todotaskapplication.data.pojo.task.TaskDto;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public interface TaskService {
    Optional<UUID> createTask(TaskDto task, String userId) throws IllegalArgumentException;

    Optional<TaskDto> getTaskById(UUID id);

    Optional<List<TaskDto>> getTasks();

    Optional<Integer> updateTask(TaskDto task, UUID taskId) throws IllegalArgumentException;

    Optional<Void> deleteTask(UUID id);
}
