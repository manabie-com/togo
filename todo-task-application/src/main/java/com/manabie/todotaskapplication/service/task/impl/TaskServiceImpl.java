package com.manabie.todotaskapplication.service.task.impl;

import com.manabie.todotaskapplication.common.constant.CustomExceptionCode;
import com.manabie.todotaskapplication.common.exception.TechnicalException;
import com.manabie.todotaskapplication.common.exception.ValidationException;
import com.manabie.todotaskapplication.common.mapper.TaskMapper;
import com.manabie.todotaskapplication.common.constant.TaskActionType;
import com.manabie.todotaskapplication.service.ratelimit.RateLimitService;
import com.manabie.todotaskapplication.service.ratelimit.impl.RateLimitServiceImpl;
import com.manabie.todotaskapplication.common.validator.TaskValidator;
import com.manabie.todotaskapplication.data.model.Task;
import com.manabie.todotaskapplication.data.pojo.task.TaskDto;
import com.manabie.todotaskapplication.repository.task.TaskRepository;
import com.manabie.todotaskapplication.service.task.TaskService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;
import org.springframework.util.CollectionUtils;

import java.time.LocalDate;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
@Service
public class TaskServiceImpl implements TaskService {
    private static final Logger LOGGER = LoggerFactory.getLogger(TaskServiceImpl.class);

    TaskRepository taskRepository;
    TaskMapper taskMapper;
    TaskValidator taskValidator;
    RateLimitService rateLimitService;
    RedisTemplate<String, Object> redisTemplate;

    @Autowired
    public TaskServiceImpl(TaskRepository taskRepository, TaskMapper taskMapper, TaskValidator taskValidator,
                           RateLimitService rateLimitService, RedisTemplate<String, Object> redisTemplate) {
        this.taskRepository = taskRepository;
        this.taskMapper = taskMapper;
        this.taskValidator = taskValidator;
        this.rateLimitService = rateLimitService;
        this.redisTemplate = redisTemplate;
    }

    @Override
    public Optional<UUID> createTask(TaskDto taskDto, String userId) throws IllegalArgumentException {
        LocalDate now = LocalDate.now();
        validateCreateTask(taskDto, userId, now);
        Task task = taskMapper.convertTaskDtoToTask(taskDto);
        task.setUserId(userId);

        if (!rateLimitService.increaseCounterAnCheckRateLimit(userId, now)) {
            throw new ValidationException(CustomExceptionCode.RATE_LIMIT_EXCEPTION, new Throwable(CustomExceptionCode.RATE_LIMIT_EXCEPTION.getValue()));
        }

        try {
            Task rs = taskRepository.save(task);
            return Optional.of(rs.getId());
        } catch (Exception e) {
            rateLimitService.decreaseCounterRateLimit(userId, now);
            LOGGER.error("Error when save task, error: {}", e.getMessage());
            throw new TechnicalException(CustomExceptionCode.INTERAL_SERVER_ERROR, new Throwable(CustomExceptionCode.INTERAL_SERVER_ERROR.getValue()));
        }
    }

    @Override
    public Optional<TaskDto> getTaskById(UUID id) {
        Optional<Task> task = taskRepository.findById(id);
        TaskDto taskDto = null;
        if (task.isPresent()) {
            taskDto = taskMapper.convertTaskToTaskDto(task.get());
        }
        return Optional.ofNullable(taskDto);
    }

    @Override
    public Optional<List<TaskDto>> getTasks() {
        List<Task> tasks = taskRepository.findAll();
        List<TaskDto> taskDtos = null;
        if (!CollectionUtils.isEmpty(tasks)) {
            taskDtos = taskMapper.convertTaskToTaskDto(tasks);
        }
        return Optional.ofNullable(taskDtos);
    }

    @Override
    public Optional<Integer> updateTask(TaskDto taskDto, UUID taskId) throws IllegalArgumentException {
        validateUpdateTask(taskDto);
        int count = taskRepository.updateTask(taskDto.getName(), taskDto.getContent(), taskId);
        return Optional.of(count);
    }

    @Override
    public Optional<Void> deleteTask(UUID id) {
        taskRepository.deleteById(id);
        return Optional.empty();
    }

    private void validateCreateTask(TaskDto taskDto, String userId, LocalDate date) {
        if (!rateLimitService.isValidRateLimit(userId, date, TaskActionType.CREATE)) {
            LOGGER.error("validate rate limit failed, too many request, userId: {}, date: {}", userId, date);
            throw new ValidationException(CustomExceptionCode.RATE_LIMIT_EXCEPTION, new Throwable(CustomExceptionCode.RATE_LIMIT_EXCEPTION.getValue()));
        }
        if (!taskValidator.validateTaskDto(taskDto, TaskActionType.CREATE)) {
            LOGGER.error("validate business rules failed, userId: {}", userId);
            throw new ValidationException(CustomExceptionCode.VALIDATE_CREATE_TASK_EXCEPTION, new Throwable(CustomExceptionCode.VALIDATE_CREATE_TASK_EXCEPTION.getValue()));
        }
    }

    private void validateUpdateTask(TaskDto taskDto) {
        if (!taskValidator.validateTaskDto(taskDto, TaskActionType.UPDATE)) {
            LOGGER.warn("validate rate limit failed, violate business rules");
            throw new ValidationException(CustomExceptionCode.VALIDATE_UPDATE_TASK_EXCEPTION, new Throwable(CustomExceptionCode.VALIDATE_UPDATE_TASK_EXCEPTION.getValue()));
        }
    }
}
