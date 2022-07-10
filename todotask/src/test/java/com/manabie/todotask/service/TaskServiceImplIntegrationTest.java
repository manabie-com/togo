package com.manabie.todotask.service;

import com.manabie.todotask.body.AddTaskRequest;
import com.manabie.todotask.entity.Task;
import com.manabie.todotask.entity.UserDailyLimit;
import com.manabie.todotask.exception.ReachLimitException;
import com.manabie.todotask.exception.UserNotFoundException;
import com.manabie.todotask.repository.TaskRepository;
import com.manabie.todotask.repository.UserRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.jdbc.AutoConfigureTestDatabase;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.annotation.DirtiesContext;

import java.time.ZonedDateTime;
import java.util.Optional;

import static org.hamcrest.CoreMatchers.is;
import static org.junit.jupiter.api.Assertions.*;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;

@SpringBootTest
@AutoConfigureTestDatabase
@DirtiesContext(classMode = DirtiesContext.ClassMode.BEFORE_EACH_TEST_METHOD)
class TaskServiceImplIntegrationTest {
    @Autowired
    TaskRepository taskRepository;

    @Autowired
    UserRepository userRepository;

    @Autowired
    TaskServiceImpl taskService;

    @BeforeEach
    public void setup(){
        taskRepository.deleteAll();
        userRepository.deleteAll();
        UserDailyLimit userDailyLimit = new UserDailyLimit();
        userDailyLimit.setUserId(1);
        userDailyLimit.setDailyTaskLimit(5);
        userRepository.save(userDailyLimit);
    }

    @Test
    void addTaskAndSuccess() {
        int userId = 1;
        ZonedDateTime targetDate = ZonedDateTime.parse("2022-07-08T00:00:00Z");

        AddTaskRequest request = new AddTaskRequest();
        request.setUserId(userId);
        request.setTaskName("test");
        request.setTaskDescription("test");
        request.setTargetDate(targetDate);

        Task task = taskService.addTask(request);

        assert(task.getId().equals(1));
        assert(task.getUserId().equals(userId));
    }

    @Test
    void addTaskWithUserNotExist() {
        int userId = 2;
        ZonedDateTime targetDate = ZonedDateTime.parse("2022-07-08T00:00:00Z");

        AddTaskRequest request = new AddTaskRequest();
        request.setUserId(userId);
        request.setTaskName("test");
        request.setTaskDescription("test");
        request.setTargetDate(targetDate);

        UserNotFoundException e = assertThrows(UserNotFoundException.class, () -> taskService.addTask(request));

        assert(e.getUserId().equals(userId));
    }

    @Test
    void addTaskWithUserReachLimitTask() {
        int userId = 1;
        Optional<UserDailyLimit> userDailyLimitOpt = userRepository.findById(userId);
        assert (userDailyLimitOpt.isPresent());
        UserDailyLimit userDailyLimit = userDailyLimitOpt.get();

        ZonedDateTime targetDate = ZonedDateTime.parse("2022-07-08T00:00:00Z");

        AddTaskRequest request = new AddTaskRequest();
        request.setUserId(userId);
        request.setTaskName("test");
        request.setTaskDescription("test");
        request.setTargetDate(targetDate);

        ReachLimitException e = assertThrows(ReachLimitException.class, () -> {
            for (int i = 0; i <= userDailyLimit.getDailyTaskLimit(); i ++) {
                taskService.addTask(request);
            }
        });
        assert(e.getDateTime().equals(request.getTargetDate()));
        assert(e.getUserId().equals(userId));
    }
}