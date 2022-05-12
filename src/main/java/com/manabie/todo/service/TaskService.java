package com.manabie.todo.service;

import com.manabie.todo.domain.Task;
import com.manabie.todo.domain.User;
import com.manabie.todo.entity.TaskEntity;
import com.manabie.todo.exception.TaskLimitException;
import com.manabie.todo.mapper.TaskMapper;
import com.manabie.todo.repository.TaskRepository;
import com.manabie.todo.utils.TimeUtils;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;

import java.text.MessageFormat;
import java.util.stream.Collectors;

@Service
@AllArgsConstructor
@Slf4j
public class TaskService {
  private final TaskRepository taskRepository;

  public Task add(User user, Task task) {
    var tasks = taskRepository.findByUserId(user.getId());
    tasks = tasks.stream().filter(x-> TimeUtils.isCreatedToday(x.getDatetimeCreated())).collect(Collectors.toList());

    log.debug("User: {}, created task:{}", user.getId(), task.toString());

    if (tasks.size() >= user.getTaskQuote())
      throw new TaskLimitException(MessageFormat.format("User: {0}, limited task:{1}", user.getId(), user.getTaskQuote()));

    var taskEntity = TaskEntity.builder()
        .userId(user.getId())
        .title(task.getTitle())
        .description(task.getDescription())
        .build();

    var saved = taskRepository.saveAndFlush(taskEntity);

    return TaskMapper.INSTANCE.toDto(saved);
  }

  public void delete(Task task) {
    log.debug("Delete task by id: {}", task.getId());
    taskRepository.deleteById(task.getId());
  }
}
