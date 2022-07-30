package com.uuhnaut69.app.todo.service;

import com.uuhnaut69.app.common.exception.MaximumLimitConfigException;
import com.uuhnaut69.app.todo.model.Todo;
import com.uuhnaut69.app.todo.model.dto.TodoRequest;
import com.uuhnaut69.app.todo.repository.TodoRepository;
import com.uuhnaut69.app.user.service.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Transactional;

/**
 * @author uuhnaut
 */
@Service
@RequiredArgsConstructor
public class TodoService {

  private final UserService userService;
  private final TodoRepository todoRepository;

  /**
   * Create new todo
   *
   * @param todoRequest Create {@link TodoRequest}
   * @return Return {@link Todo}
   */
  @Transactional(isolation = Isolation.SERIALIZABLE, rollbackFor = Exception.class)
  public Todo createNewTodo(TodoRequest todoRequest) {
    // Get creation limit of user
    var user = userService.findUserById(todoRequest.getUserId());
    var limit = user.getLimitConfig();

    // Get number of created todos today
    var numberOfCreatedTodosByUserId = todoRepository.countNumberOfCreatedTodosTodayByUserId(
        user.getId());

    if (numberOfCreatedTodosByUserId >= limit) {
      throw new MaximumLimitConfigException("Reach maximum creation todo per day");
    }

    var todo = new Todo();
    todo.setTask(todoRequest.getTask());
    todo.setUserId(todoRequest.getUserId());

    return todoRepository.save(todo);
  }
}
