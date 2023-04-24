package com.manabie.todo.service.impl;

import com.manabie.todo.constant.TaskStatus;
import com.manabie.todo.entity.TaskEntity;
import com.manabie.todo.exception.ManabieException;
import com.manabie.todo.model.CreateTaskRequest;
import com.manabie.todo.model.TaskModel;
import com.manabie.todo.model.UserInfo;
import com.manabie.todo.repository.TaskRepository;
import com.manabie.todo.service.TaskLimitService;
import com.manabie.todo.service.TaskService;
import lombok.RequiredArgsConstructor;
import org.modelmapper.ModelMapper;
import org.springframework.http.HttpStatus;
import org.springframework.integration.redis.util.RedisLockRegistry;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;

import java.time.LocalDateTime;
import java.util.List;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.Lock;

@Service
@RequiredArgsConstructor
public class TaskServiceImpl implements TaskService {
    final ModelMapper modelMapper;
    final TaskRepository taskRepository;
    final TaskLimitService taskLimitService;
    final RedisLockRegistry userLockRegistry;
    final WebClient.Builder webClientBuilder;

    @Override
    public TaskModel createTask(CreateTaskRequest request) {
        UserInfo userInfo = getUserInfo(request.getOwner());
        if (null == userInfo) {
            throw new ManabieException(HttpStatus.NOT_FOUND.ordinal(), "User not found", HttpStatus.NOT_FOUND);
        }

        TaskModel taskModel = null;
        Lock lock = userLockRegistry.obtain(request.getOwner() + "");
        boolean isLock = false;
        try {
            isLock = lock.tryLock(100, TimeUnit.MILLISECONDS);
            if (isLock) {
                if (!taskLimitService.checkLimitAndIncreaseCounter(request.getOwner(),userInfo.getMaxTaskPerDay())) {
                    TaskEntity taskEntity = modelMapper.map(request, TaskEntity.class);
                    taskEntity.setStatus(TaskStatus.NEW);
                    taskEntity.setCreatedAt(LocalDateTime.now());

                    taskModel = modelMapper.map(taskRepository.save(taskEntity), TaskModel.class);
                }
                lock.unlock();
            }
        } catch (Exception e) {
            if (isLock) {
                lock.unlock();
            }
        }
        return taskModel;
    }

    @Override
    public List<TaskModel> findAllTask() {
        return taskRepository.findAll().stream().map(t -> modelMapper.map(t, TaskModel.class)).toList();
    }

    @Override
    public List<TaskModel> findAllByOwner(final Long userId) {
        return taskRepository.findAllByOwner(userId).stream().map(t -> modelMapper.map(t, TaskModel.class)).toList();
    }

    @Override
    public TaskModel getTaskById(Long id) {
        return modelMapper.map(taskRepository.getReferenceById(id), TaskModel.class);
    }

    @Override
    public UserInfo getUserInfo(Long userId) {
        return webClientBuilder.build().get()
                .uri("http://user-service:8080",
                        uriBuilder -> uriBuilder.pathSegment("user",userId + "").build())
                .retrieve()
                .bodyToMono(UserInfo.class)
                .block();
    }
}
