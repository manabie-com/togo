package com.uuhnaut69.app.todo.service;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.ArgumentMatchers.any;

import com.uuhnaut69.app.common.exception.MaximumLimitConfigException;
import com.uuhnaut69.app.common.exception.NotFoundException;
import com.uuhnaut69.app.todo.model.Todo;
import com.uuhnaut69.app.todo.model.dto.TodoRequest;
import com.uuhnaut69.app.todo.repository.TodoRepository;
import com.uuhnaut69.app.user.model.User;
import com.uuhnaut69.app.user.service.UserService;
import java.time.Instant;
import java.util.Collections;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.junit.jupiter.MockitoExtension;

/**
 * @author uuhnaut
 */
@ExtendWith(MockitoExtension.class)
class TodoServiceTest {

  @Mock
  private UserService userService;
  @Mock
  private TodoRepository todoRepository;

  @InjectMocks
  private TodoService todoService;

  @Test
  void createNewTodoShouldSuccess() {
    var todoRequest = new TodoRequest("Test", 1L);

    var user = new User(1L, "uuhnaut69", 10, Instant.now());
    var todo = new Todo(1L, "Test", 1L, Instant.now());

    Mockito.when(userService.findUserById(1L)).thenReturn(user);
    Mockito.when(todoRepository.countNumberOfCreatedTodosTodayByUserId(1L)).thenReturn(1L);
    Mockito.when(todoRepository.save(any())).thenReturn(todo);

    var result = todoService.createNewTodo(todoRequest);

    assertEquals(1L, result.getId());
    assertEquals("Test", result.getTask());
  }

  @Test
  void createNewTodoFailedByUserNotFound() {
    var todoRequest = new TodoRequest("Test", 1L);

    Mockito.when(userService.findUserById(1L)).thenThrow(new NotFoundException());

    try {
      todoService.createNewTodo(todoRequest);
    } catch (Exception e) {
      assertTrue(e instanceof NotFoundException);
    }
  }

  @Test
  void createNewTodoFailedByReachMaximumLimitConfig() {
    var todoRequest = new TodoRequest("Test", 1L);
    var user = new User(1L, "uuhnaut69", 10, Instant.now());

    Mockito.when(userService.findUserById(1L)).thenReturn(user);
    Mockito.when(todoRepository.countNumberOfCreatedTodosTodayByUserId(1L)).thenReturn(11L);

    try {
      todoService.createNewTodo(todoRequest);
    } catch (Exception e) {
      assertTrue(e instanceof MaximumLimitConfigException);
    }
  }

  @Test
  void findAllTodos() {
    var todo = new Todo(1L, "Test", 1L, Instant.now());
    Mockito.when(todoRepository.findAll()).thenReturn(Collections.singletonList(todo));

    var result = todoService.findAllTodo(null);
    assertEquals(1, result.size());
  }

  @Test
  void findAllTodosByUserId() {
    var todo = new Todo(1L, "Test", 1L, Instant.now());
    Mockito.when(todoRepository.findAllByUserIdOrderByIdDesc(1L))
        .thenReturn(Collections.singletonList(todo));

    var result = todoService.findAllTodo(1L);
    assertEquals(1, result.size());
  }
}