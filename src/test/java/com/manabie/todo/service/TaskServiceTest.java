package com.manabie.todo.service;

import com.manabie.todo.domain.Task;
import com.manabie.todo.domain.User;
import com.manabie.todo.entity.TaskEntity;
import com.manabie.todo.entity.UserEntity;
import com.manabie.todo.exception.TaskLimitException;
import com.manabie.todo.repository.TaskRepository;
import com.manabie.todo.repository.UserRepository;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.invocation.InvocationOnMock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.mockito.stubbing.Answer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Profile;
import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.test.context.ActiveProfiles;

import javax.persistence.Table;
import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ActiveProfiles(profiles = "test")
@SpringBootTest
@EnableAutoConfiguration
class TaskServiceTest {

  @Autowired
  TaskService taskService;

  @Autowired
  TaskRepository taskRepository;

  @Autowired
  UserRepository userRepository;

  User user;
  UserEntity userEntity;
  Task task;
  TaskEntity taskEntity;

  @BeforeEach
  void setUp() {
    user = User.builder()
        .id(1L)
        .username("long")
        .password("password")
        .taskQuote(2L)
        .build();
    userEntity = UserEntity.builder()
        .id(1L)
        .username("long")
        .password("password")
        .taskQuote(2L)
        .build();
    task = Task.builder()
        .title("test")
        .description("test description")
        .userId(1L)
        .build();
    taskEntity = TaskEntity.builder()
        .title("test")
        .description("test description")
        .userId(1L)
        .build();

    userRepository.save(userEntity);
  }


  @AfterEach
  void tearDown() {
    taskRepository.deleteAll();
  }

  @Test
  void giveValidTask_whenAdd_thenSucceed() {
    taskRepository.save(taskEntity);

    var result = taskService.add(user,task);

    assertEquals(taskRepository.findByUserId(user.getId()).size(),2);
    assertNotNull(result.getId());
  }

  @Test
  void giveLimitTask_whenAdd_thenFailed() {
    taskRepository.save(taskEntity);
    user.setTaskQuote(1L);

    assertThrows(TaskLimitException.class,()->{
      taskService.add(user,task);
    });
  }

  @Test
  void giveExistTask_whenDelete_thenSucceed() {
    taskRepository.save(taskEntity);
    var taskInsert = taskRepository.findAll().get(0);

    task.setId(taskInsert.getId());
    long before = taskRepository.count();
    taskService.delete(task);
    long after = taskRepository.count();

    assertEquals(before-1,after);
  }

  @Test
  void giveNonExistTask_whenDelete_thenFailed() {
    taskRepository.save(taskEntity);
    var taskInsert = taskRepository.findAll().get(0);

    task.setId(taskInsert.getId()+1);

    assertThrows(EmptyResultDataAccessException.class,()->{
      taskService.delete(task);
    });
  }
}