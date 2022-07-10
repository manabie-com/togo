package com.manabie.todotask.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.manabie.todotask.body.AddTaskRequest;
import com.manabie.todotask.entity.Task;
import com.manabie.todotask.entity.UserDailyLimit;
import com.manabie.todotask.repository.TaskRepository;
import com.manabie.todotask.repository.UserRepository;
import com.manabie.todotask.service.TaskService;
import javafx.application.Application;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.security.servlet.SecurityAutoConfiguration;
import org.springframework.boot.test.autoconfigure.jdbc.AutoConfigureTestDatabase;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.mock.web.MockHttpServletRequest;
import org.springframework.test.annotation.DirtiesContext;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.RequestBuilder;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.Matchers.greaterThanOrEqualTo;
import static org.hamcrest.Matchers.hasSize;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import java.time.ZonedDateTime;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest
@AutoConfigureMockMvc
@EnableAutoConfiguration(exclude= SecurityAutoConfiguration.class)
@AutoConfigureTestDatabase
@DirtiesContext(classMode = DirtiesContext.ClassMode.BEFORE_EACH_TEST_METHOD)
class TaskControllerIntegrationTest {
    @Autowired
    MockMvc mockMvc;

    @Autowired
    TaskService taskService;
    @Autowired
    UserRepository userRepository;
    @Autowired
    TaskRepository taskRepository;

    ObjectMapper mapper = new ObjectMapper().findAndRegisterModules();

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
    void addTaskAndReturnSuccess() throws Exception {
        int userId = 1;
        ZonedDateTime targetDate = ZonedDateTime.parse("2022-07-08T00:00:00Z");
        AddTaskRequest request = new AddTaskRequest();
        request.setUserId(userId);
        request.setTaskName("test");
        request.setTaskDescription("test");
        request.setTargetDate(targetDate);

        RequestBuilder requestBuilder = MockMvcRequestBuilders
                .put("/task/add")
                .content(mapper.writeValueAsBytes(request))
                .contentType(MediaType.APPLICATION_JSON);

        mockMvc.perform(requestBuilder)
                .andDo(print())
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code", is(0)))
                .andExpect(jsonPath("$.data.id", is(1)));
    }

    @Test
    void addTaskWithUserNotExist() throws Exception {
        int userId = 2;
        ZonedDateTime targetDate = ZonedDateTime.parse("2022-07-08T00:00:00Z");
        AddTaskRequest request = new AddTaskRequest();
        request.setUserId(userId);
        request.setTaskName("test");
        request.setTaskDescription("test");
        request.setTargetDate(targetDate);

        RequestBuilder requestBuilder = MockMvcRequestBuilders
                .put("/task/add")
                .content(mapper.writeValueAsBytes(request))
                .contentType(MediaType.APPLICATION_JSON);

        mockMvc.perform(requestBuilder)
                .andDo(print())
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("$.code", is(1)))
                .andExpect(jsonPath("$.message", is("User with id " + userId + " does not exist")));
    }

    @Test
    void addTaskWithReachLimitTask() throws Exception {
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

        RequestBuilder requestBuilder = MockMvcRequestBuilders
                .put("/task/add")
                .content(mapper.writeValueAsBytes(request))
                .contentType(MediaType.APPLICATION_JSON);

        //add until reach limit
        for (int i = 0; i < userDailyLimit.getDailyTaskLimit(); i ++) {
            mockMvc.perform(requestBuilder)
                    .andDo(print())
                    .andExpect(status().isOk())
                    .andExpect(jsonPath("$.code", is(0)))
                    .andExpect(jsonPath("$.data.user_id", is(userId)));
        }
        //add when limit is reached
        mockMvc.perform(requestBuilder)
                .andDo(print())
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("$.code", is(1)));
    }
}