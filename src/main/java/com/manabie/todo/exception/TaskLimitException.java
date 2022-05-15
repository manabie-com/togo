package com.manabie.todo.exception;

import com.manabie.todo.utils.AppCode;
import lombok.Getter;

@Getter
public class TaskLimitException extends AppException {

  private String message;

  public TaskLimitException() {
    super(AppCode.TASK_LIMIT, "");
  }

  public TaskLimitException(String message) {
    super(AppCode.TASK_LIMIT, message);
  }
}
