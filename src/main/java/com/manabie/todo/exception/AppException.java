package com.manabie.todo.exception;

import com.manabie.todo.utils.AppCode;
import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public class AppException extends RuntimeException {
  private final AppCode code;
  private String message;

  public AppException(AppCode code, Throwable throwable) {
    super(throwable);
    this.code = code;
  }
}
