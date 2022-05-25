package com.todo.core.todo.service;

import com.todo.core.commons.Messages;
import com.todo.core.commons.model.GenericResponse;
import com.todo.core.todo.application.dto.TodoDTO;
import com.todo.core.todo.exception.LimitForTodayReachedException;
import com.todo.core.todo.model.Todo;
import com.todo.core.todo.repository.TodoRepository;
import com.todo.core.user.model.TodoUser;
import com.todo.core.user.repository.TodoUserRepository;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDate;
import java.time.ZoneId;


@Service
@Transactional
public class TodoServiceImpl implements TodoService {

    private TodoRepository todoRepository;
    private TodoUserRepository todoUserRepository;

    public TodoServiceImpl(TodoRepository todoRepository, TodoUserRepository todoUserRepository) {
        this.todoRepository = todoRepository;
        this.todoUserRepository = todoUserRepository;
    }

    @Override
    @Transactional(readOnly = false)
    public GenericResponse<Integer> addTodo(Long todoUserId, TodoDTO todoDto) {
        final Todo todoForSave = new Todo(todoDto, todoUserId, LocalDate.now());

        if (getAddableToday(todoUserId) > 0) {
            todoRepository.save(todoForSave);
            return new GenericResponse<>(
                1,
                Messages.SAVE_SUCCESSFUL.getContent()
            );
        } else {
            throw new LimitForTodayReachedException("User has exceeded limit for today!");
        }

    }

    @Override
    @Transactional(readOnly = true)
    public GenericResponse<Page<Todo>> retrieveAllTodoByUser(Long userId, Pageable pageable) {
        final Page<Todo> pagedTodos = this.todoRepository.findAllByTodoUserId(userId, pageable);

        return new GenericResponse<>(
            pagedTodos,
            Messages.PAGES_RETRIEVE_SUCCESSFUL.getContent()
        );
    }

    private int getAddableToday(Long userId) {
        final TodoUser todoUser = this.todoUserRepository
            .findById(userId).orElseThrow(() -> new RuntimeException("User not found with id: " + userId));

        final int countForDate = this.todoRepository
            .countByDateCreatedAndTodoUserId(LocalDate.now(ZoneId.of("Asia/Manila")), userId);

        return (int) (todoUser.getTodoLimit() - countForDate);
    }
}
