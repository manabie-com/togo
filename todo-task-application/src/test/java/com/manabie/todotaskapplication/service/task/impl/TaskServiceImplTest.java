package com.manabie.todotaskapplication.service.task.impl;

import com.manabie.todotaskapplication.common.exception.TechnicalException;
import com.manabie.todotaskapplication.common.exception.ValidationException;
import com.manabie.todotaskapplication.common.mapper.TaskMapper;
import com.manabie.todotaskapplication.service.ratelimit.impl.RateLimitServiceImpl;
import com.manabie.todotaskapplication.common.validator.TaskValidator;
import com.manabie.todotaskapplication.data.model.Task;
import com.manabie.todotaskapplication.data.pojo.task.TaskDto;
import com.manabie.todotaskapplication.repository.task.TaskRepository;
import com.manabie.todotaskapplication.service.userconfig.UserConfigService;
import org.junit.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.http.HttpStatus;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.*;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
@RunWith(SpringRunner.class)
public class TaskServiceImplTest {
    @Mock
    TaskRepository taskRepository;
    @Mock
    TaskMapper taskMapper;
    @Mock
    TaskValidator taskValidator;
    @Mock
    RateLimitServiceImpl rateLimitValidator;

    @Mock
    private UserConfigService userConfigService;
    @InjectMocks
    TaskServiceImpl taskService;

    @Test
    public void createTask_Success() {
        int rateLimit = 3;
        TaskDto taskDto = getValidTaskDto();
        String userId = "123213";
        UUID id = UUID.randomUUID();
        Task task = new Task(id, taskDto.getName(), taskDto.getContent(), userId);
        when(rateLimitValidator.isValidRateLimit(any(), any(), any())).thenReturn(true);
        when(taskValidator.validateTaskDto(any(), any())).thenReturn(true);
        when(taskMapper.convertTaskDtoToTask(taskDto)).thenReturn(task);
        when(rateLimitValidator.increaseCounterAnCheckRateLimit(any(), any())).thenReturn(true);
        when(taskRepository.save(any())).thenReturn(task);
        Optional<UUID> taskId = taskService.createTask(taskDto, userId);
        assertTrue(taskId.isPresent());
        assertEquals(id, taskId.get());
    }

    @Test
    public void createTask_Failed_User_Rate_Limit_Exceed() {
        TaskDto taskDto = getValidTaskDto();
        String userId = "123213";
        UUID id = UUID.randomUUID();
        Task task = new Task(id, taskDto.getName(), taskDto.getContent(), userId);
        when(rateLimitValidator.isValidRateLimit(any(), any(), any())).thenReturn(false);
        ValidationException validationException = assertThrows(ValidationException.class, () -> taskService.createTask(taskDto, userId));
        assertEquals(HttpStatus.TOO_MANY_REQUESTS, validationException.getStatus());
    }

    @Test
    public void createTask_Failed_Validate_Business_Rules() {
        int rateLimit = 3;
        TaskDto taskDto = getInvalidTaskDto();
        String userId = "123213";
        UUID id = UUID.randomUUID();
        Task task = new Task(id, taskDto.getName(), taskDto.getContent(), userId);
        when(rateLimitValidator.isValidRateLimit(any(), any(), any())).thenReturn(true);
        when(taskValidator.validateTaskDto(any(), any())).thenReturn(false);
        ValidationException validationException = assertThrows(ValidationException.class, () -> taskService.createTask(taskDto, userId));
        assertEquals(HttpStatus.BAD_REQUEST, validationException.getStatus());
    }

    @Test
    public void createTask_Failed_Rate_Limit_When_Increment_And_Check_Counter() {
        int rateLimit = 3;
        TaskDto taskDto = getInvalidTaskDto();
        String userId = "123213";
        UUID id = UUID.randomUUID();
        Task task = new Task(id, taskDto.getName(), taskDto.getContent(), userId);
        when(rateLimitValidator.isValidRateLimit(any(), any(), any())).thenReturn(true);
        when(taskValidator.validateTaskDto(any(), any())).thenReturn(true);
        when(taskMapper.convertTaskDtoToTask(taskDto)).thenReturn(task);
        when(rateLimitValidator.increaseCounterAnCheckRateLimit(any(), any())).thenReturn(false);
        ValidationException validationException = assertThrows(ValidationException.class, () -> taskService.createTask(taskDto, userId));
        assertEquals(HttpStatus.TOO_MANY_REQUESTS, validationException.getStatus());
    }

    @Test
    public void createTask_Failed_Save_Task_Throw_Exception() {
        int rateLimit = 3;
        TaskDto taskDto = getInvalidTaskDto();
        String userId = "123213";
        UUID id = UUID.randomUUID();
        Task task = new Task(id, taskDto.getName(), taskDto.getContent(), userId);
        when(rateLimitValidator.isValidRateLimit(any(), any(), any())).thenReturn(true);
        when(taskValidator.validateTaskDto(any(), any())).thenReturn(true);
        when(taskMapper.convertTaskDtoToTask(taskDto)).thenReturn(task);
        when(rateLimitValidator.increaseCounterAnCheckRateLimit(any(), any())).thenReturn(true);
        when(taskRepository.save(any())).thenThrow(IllegalArgumentException.class);
        when(rateLimitValidator.decreaseCounterRateLimit(any(), any())).thenReturn(1L);
        TechnicalException technicalException = assertThrows(TechnicalException.class, () -> taskService.createTask(taskDto, userId));
        assertEquals(HttpStatus.INTERNAL_SERVER_ERROR, technicalException.getStatus());
    }

    @Test
    public void getTaskById_Success() {
        TaskDto taskDto = getInvalidTaskDto();
        String userId = "123213";
        UUID id = UUID.randomUUID();
        Task task = new Task(id, taskDto.getName(), taskDto.getContent(), userId);
        when(taskRepository.findById(id)).thenReturn(Optional.of(task));
        when(taskMapper.convertTaskToTaskDto(task)).thenReturn(new TaskDto(task.getId(), task.getName(), task.getContent(), task.getUserId()));
        Optional<TaskDto> rs = taskService.getTaskById(id);
        assertTrue(rs.isPresent());
        assertEquals(id, rs.get().getId());
    }

    @Test
    public void getTaskById_Failed_Not_Found() {
        TaskDto taskDto = getValidTaskDto();
        UUID id = UUID.randomUUID();
        Optional<TaskDto> rs = taskService.getTaskById(id);
        assertFalse(rs.isPresent());
    }

    @Test
    public void getTasks_Success() {
        List<Task> tasks = Arrays.asList(new Task(UUID.randomUUID(), "test-1", "content-test-1", "123123"), new Task(UUID.randomUUID(), "test-1", "content-test-1", "123123"));
        when(taskRepository.findAll()).thenReturn(tasks);
        when(taskMapper.convertTaskToTaskDto(tasks)).thenReturn(Arrays.asList(new TaskDto(tasks.get(0).getId(), tasks.get(0).getName(), tasks.get(0).getContent(), tasks.get(0).getUserId()),
                new TaskDto(tasks.get(1).getId(), tasks.get(1).getName(), tasks.get(1).getContent(), tasks.get(1).getUserId())));
        Optional<List<TaskDto>> taskDtos = taskService.getTasks();
        assertTrue(taskDtos.isPresent());
        assertTrue(taskDtos.get().size() > 0);
    }

    @Test
    public void getTasks_Not_Found() {
        when(taskRepository.findAll()).thenReturn(Collections.emptyList());
        Optional<List<TaskDto>> taskDtos = taskService.getTasks();
        assertFalse(taskDtos.isPresent());
    }

    @Test
    public void updateTask_Success() {
        TaskDto taskDto = getValidTaskDto();
        UUID taskId = UUID.randomUUID();
        when(taskValidator.validateTaskDto(any(), any())).thenReturn(true);
        when(taskRepository.updateTask(any(), any(), any())).thenReturn(1);
        Optional<Integer> updateCounter = taskService.updateTask(taskDto, taskId);
        assertTrue(updateCounter.isPresent());
        assertEquals(1, (int) updateCounter.get());
    }

    @Test
    public void updateTask_Failed_Not_Found() {
        TaskDto taskDto = getValidTaskDto();
        UUID taskId = UUID.randomUUID();
        when(taskValidator.validateTaskDto(any(), any())).thenReturn(true);
        when(taskRepository.updateTask(any(), any(), any())).thenReturn(0);
        Optional<Integer> updateCounter = taskService.updateTask(taskDto, taskId);
        assertTrue(updateCounter.isPresent());
        assertEquals(0, (int) updateCounter.get());
    }

    @Test
    public void updateTask_Failed_Validate_Business_Rules() {
        TaskDto taskDto = getValidTaskDto();
        UUID taskId = UUID.randomUUID();
        when(taskValidator.validateTaskDto(any(), any())).thenReturn(false);
        ValidationException validationException = assertThrows(ValidationException.class, () -> taskService.updateTask(any(), any()));
        assertEquals(HttpStatus.BAD_REQUEST, validationException.getStatus());
    }

    @Test
    public void deleteTask() {
        UUID taskID = UUID.randomUUID();
        taskService.deleteTask(taskID);
        assertTrue(true);
    }

    private TaskDto getValidTaskDto() {
        TaskDto taskDto = new TaskDto();
        taskDto.setName("test-1");
        taskDto.setContent("content-test-1");
        return taskDto;
    }

    private TaskDto getInvalidTaskDto() {
        TaskDto taskDto = new TaskDto();
        taskDto.setContent("content-test-1");
        return taskDto;
    }
}