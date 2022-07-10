package com.manabie.todotask.service;

import com.manabie.todotask.body.AddTaskRequest;
import com.manabie.todotask.entity.Task;
import com.manabie.todotask.exception.ReachLimitException;
import com.manabie.todotask.exception.UserNotFoundException;
import com.manabie.todotask.repository.TaskRepository;
import com.manabie.todotask.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.ZonedDateTime;

@Service
public class TaskServiceImpl implements TaskService{
    @Autowired
    TaskRepository taskRepository;
    @Autowired
    UserRepository userRepository;

    public Task addTask(AddTaskRequest request){
        ZonedDateTime startTime = request.getTargetDate().toLocalDate().atStartOfDay().atZone(request.getTargetDate().getZone());
        ZonedDateTime endTime = request.getTargetDate().toLocalDate().atStartOfDay().plusDays(1).atZone(request.getTargetDate().getZone());
        return userRepository.findById(request.getUserId()).map(userDailyLimit -> {
            int taskCount = taskRepository.countTaskByUserIdAndTargetDate(request.getUserId(), startTime, endTime);
            if (taskCount < userDailyLimit.getDailyTaskLimit()) {
                return taskRepository.save(buildFromRequest(request));
            } else {
                throw new ReachLimitException(request.getUserId(), request.getTargetDate());
            }
        }).orElseThrow(() -> new UserNotFoundException(request.getUserId()));
    }

    private Task buildFromRequest(AddTaskRequest request){
        Task task = new Task();
        task.setUserId(request.getUserId());
        task.setDescription(request.getTaskDescription());
        task.setName(request.getTaskName());
        task.setTargetDate(request.getTargetDate());
        return task;
    }
}
