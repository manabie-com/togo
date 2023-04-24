package com.manabie.todo.controller;

import com.manabie.todo.model.CreateTaskRequest;
import com.manabie.todo.model.TaskModel;
import com.manabie.todo.model.UserInfo;
import com.manabie.todo.repository.TaskRepository;
import com.manabie.todo.service.TaskLimitService;
import com.manabie.todo.service.TaskService;
import com.manabie.todo.service.impl.TaskServiceImpl;
import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import org.mockito.MockitoAnnotations;
import org.modelmapper.ModelMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.integration.redis.util.RedisLockRegistry;
import org.springframework.web.reactive.function.client.WebClient;

import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.when;

@Slf4j
public class TaskServiceTest {

    @Mock
    private ModelMapper modelMapper;
    @Mock
    private TaskRepository taskRepository;
    @Mock
    private TaskLimitService taskLimitService;
    @Autowired
    private RedisLockRegistry userLockRegistry;
    @Mock
    private WebClient.Builder webClientBuilder;
    @Mock
    private TaskService taskService;


    @BeforeEach
    public void setUp() {
        MockitoAnnotations.openMocks(this);
        taskService = new TaskServiceImpl(modelMapper,taskRepository,taskLimitService,userLockRegistry,webClientBuilder);
    }


    @Test
    public void createTask_returnSuccess() throws Exception {
        CreateTaskRequest request=new CreateTaskRequest();
        request.setTitle("Task 1");
        request.setDescription("This is Task 1");
        request.setOwner(1000L);

        UserInfo fakedUserInfo=new UserInfo();
        fakedUserInfo.setId(1000L);
        fakedUserInfo.setUsername("trungnguyen.tech");
        fakedUserInfo.setId(1000L);
        fakedUserInfo.setMaxTaskPerDay(3);

        when(taskService.getUserInfo(any())).thenReturn(fakedUserInfo);

        TaskModel taskModel = taskService.createTask(request);
        assertNotNull(taskModel);
    }
}
