package com.manabie.todo.api.request;

import com.manabie.todo.utils.Constants;
import lombok.Data;

import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.Pattern;

@Data
public class UpdateUserRequest {
  private String username;

  @Pattern(
      regexp = Constants.PASSWORD_SECURITY_PATTERN,
      message = "{app.login.request.password.security.require}")
  private String password;

  private Long taskQuote;
}
