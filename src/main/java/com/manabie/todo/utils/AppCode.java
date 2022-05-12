package com.manabie.todo.utils;

public enum AppCode {
  SUCCESS("code.success"),
  USER_ERROR("code.user.error"),
  USER_DUPLICATED("code.user.duplicated"),
  USER_UNAUTHORIZED("code.user.unauthorized"),
  TASK_LIMIT("code.task.limited");
  final String message;

  AppCode(String message) {
    this.message = message;
  }

  public String getMessage() {
    return message;
  }
}
