package com.manabie.todotaskapplication.common.mapper;

import com.manabie.todotaskapplication.data.model.Task;
import com.manabie.todotaskapplication.data.pojo.task.TaskDto;

import java.util.List;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public interface TaskMapper {
    Task convertTaskDtoToTask(TaskDto taskDto);

    TaskDto convertTaskToTaskDto(Task task);
    List<TaskDto> convertTaskToTaskDto(List<Task> tasks);

    Task mergeTaskDtoToTask(TaskDto taskDto, Task task);
}
