package com.manabie.todotask.service;

import com.manabie.todotask.body.AddTaskRequest;
import com.manabie.todotask.entity.Task;
import com.manabie.todotask.exception.ReachLimitException;
import com.manabie.todotask.exception.UserNotFoundException;
import com.manabie.todotask.repository.TaskRepository;
import com.manabie.todotask.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.client.HttpClientErrorException;

import javax.management.BadAttributeValueExpException;
import java.time.LocalDate;
import java.time.ZonedDateTime;

@Service
public class TaskServiceImpl implements TaskService{
    @Autowired
    TaskRepository taskRepository;
    @Autowired
    UserRepository userRepository;

    public Task addTask(AddTaskRequest request){
        return userRepository.findById(request.getUserId()).map(userDailyLimit -> {
            int taskCount = taskRepository.countTaskByUserIdAndTargetDate(request.getUserId(), request.getTargetDate(), request.getTargetDate().plusDays(1));
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
