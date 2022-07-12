package com.todo.ws.todo.service;

import com.todo.ws.todo.application.dto.TodoDTO;
import com.todo.ws.todo.exception.LimitForTodayReachedException;
import com.todo.ws.todo.model.Todo;
import com.todo.ws.todo.repository.TodoRepository;
import com.todo.ws.user.model.TodoUser;
import com.todo.ws.user.repository.TodoUserRepository;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;


import java.time.LocalDate;
import java.util.Optional;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
public class TodoServiceTests {

    @InjectMocks
    private TodoServiceImpl todoService;

    @Mock
    private TodoRepository todoRepository;

    @Mock
    private TodoUserRepository todoUserRepository;


    @Test
    public void addTodoSuccess() {

        final TodoDTO mockTodo = new TodoDTO();
        final Optional<TodoUser> todoUser = Optional.of(mock(TodoUser.class));

        when(todoUserRepository.findById(1L)).thenReturn(todoUser);
        when(todoUser.get().getTodoLimit()).thenReturn(5L);

        when(todoRepository.countByDateCreatedAndTodoUserId(LocalDate.now(), 1L)).thenReturn(1);
        when(todoRepository.save(any(Todo.class))).thenReturn(new Todo());

        assertThat(todoService.addTodo(1L, mockTodo).getData()).isEqualTo(1);

    }

    @Test()
    public void addTodoLimitExceeded() {

        final TodoDTO mockTodo = new TodoDTO();
        final Optional<TodoUser> todoUser = Optional.of(mock(TodoUser.class));

        when(todoUserRepository.findById(1L)).thenReturn(todoUser);
        when(todoUser.get().getTodoLimit()).thenReturn(5L);
        when(todoRepository.countByDateCreatedAndTodoUserId(LocalDate.now(), 1L)).thenReturn(5);

        Assertions.assertThrows(LimitForTodayReachedException.class, () -> todoService.addTodo(1L, mockTodo));

    }
}
