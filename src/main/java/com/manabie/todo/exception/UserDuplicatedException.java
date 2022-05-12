package com.manabie.todo.exception;

import com.manabie.todo.utils.AppCode;

public class UserDuplicatedException extends AppException {

  public UserDuplicatedException() {
    super(AppCode.USER_DUPLICATED, "");
  }
}
