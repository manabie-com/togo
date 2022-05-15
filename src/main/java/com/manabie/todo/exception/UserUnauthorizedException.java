package com.manabie.todo.exception;

import com.manabie.todo.utils.AppCode;

public class UserUnauthorizedException extends AppException {

  public UserUnauthorizedException() {
    super(AppCode.USER_UNAUTHORIZED, "");
  }
}
