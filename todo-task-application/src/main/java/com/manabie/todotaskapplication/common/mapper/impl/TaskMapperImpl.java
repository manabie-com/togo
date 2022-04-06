package com.manabie.todotaskapplication.common.mapper.impl;

import com.manabie.todotaskapplication.common.mapper.TaskMapper;
import com.manabie.todotaskapplication.data.model.Task;
import com.manabie.todotaskapplication.data.pojo.task.TaskDto;
import org.springframework.stereotype.Component;
import org.springframework.util.CollectionUtils;

import java.util.Collections;
import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
@Component
public class TaskMapperImpl implements TaskMapper {
    @Override
    public Task convertTaskDtoToTask(TaskDto taskDto) {
        if (Objects.nonNull(taskDto)) {
            return Task.builder().id(taskDto.getId())
                    .name(taskDto.getName())
                    .content(taskDto.getContent())
                    .build();
        }
        return null;
    }

    @Override
    public TaskDto convertTaskToTaskDto(Task task) {
        if (Objects.nonNull(task)) {
            return TaskDto.builder().id(task.getId())
                    .name(task.getName())
                    .content(task.getContent())
                    .userId(task.getUserId())
                    .build();
        }
        return null;
    }

    @Override
    public List<TaskDto> convertTaskToTaskDto(List<Task> tasks) {
        if (CollectionUtils.isEmpty(tasks)) {
            return Collections.emptyList();
        }
        return tasks.stream().map(this::convertTaskToTaskDto).collect(Collectors.toList());
    }

    @Override
    public Task mergeTaskDtoToTask(TaskDto taskDto, Task task) {
        if (Objects.isNull(task)) {
            return null;
        }
        if (Objects.nonNull(taskDto.getName())) {
            task.setName(taskDto.getName());
        }
        if (Objects.nonNull(taskDto.getContent())) {
            task.setContent(taskDto.getName());
        }
        return task;
    }
}
