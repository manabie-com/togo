package com.uuhnaut69.app.todo.api;

import com.uuhnaut69.app.todo.model.Todo;
import com.uuhnaut69.app.todo.model.dto.TodoRequest;
import com.uuhnaut69.app.todo.service.TodoService;
import java.util.List;
import javax.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author uuhnaut
 */
@RestController
@RequiredArgsConstructor
@RequestMapping("/todos")
public class TodoController {

  private final TodoService todoService;

  /**
   * Create new todo
   *
   * @param todoRequest Create new {@link TodoRequest}
   * @return Return {@link Todo}
   */
  @PostMapping
  @ResponseStatus(HttpStatus.CREATED)
  public Todo createNewTodo(@RequestBody @Valid TodoRequest todoRequest) {
    return todoService.createNewTodo(todoRequest);
  }

  /**
   * Find todos
   *
   * @param userId Filter by user id
   * @return return {@link List} contains all {@link Todo}
   */
  @GetMapping
  public List<Todo> findAllTodo(@RequestParam(required = false) Long userId) {
    return todoService.findAllTodo(userId);
  }
}
