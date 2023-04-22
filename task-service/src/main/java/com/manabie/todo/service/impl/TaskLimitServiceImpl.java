package com.manabie.todo.service.impl;

import com.manabie.todo.entity.TaskCounterEntity;
import com.manabie.todo.repository.TaskCounterRepository;
import com.manabie.todo.service.TaskLimitService;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;

@Service
@RequiredArgsConstructor
public class TaskLimitServiceImpl implements TaskLimitService {
    final TaskCounterRepository taskCounterRepository;

    @Override
    public boolean checkLimitAndIncreaseCounter(Long userId, Integer maxTaxPerDay) {
        TaskCounterEntity taskCounter = taskCounterRepository.findByUserId(userId);
        if (null == taskCounter) {
            taskCounter = new TaskCounterEntity();
            taskCounter.setUserId(userId);
            taskCounter.setCounter(0L);
            taskCounter.setCreatedAt(LocalDateTime.now());
        }
        if (taskCounter.getCounter() < maxTaxPerDay) {
            taskCounter.setCounter(taskCounter.getCounter() + 1);
            taskCounterRepository.save(taskCounter);
            return false;
        }
        return true;
    }
}
