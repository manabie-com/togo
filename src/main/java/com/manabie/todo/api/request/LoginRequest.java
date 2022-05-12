package com.manabie.todo.api.request;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.validation.constraints.NotEmpty;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class LoginRequest {
  @NotEmpty(message = "{app.login.request.username.require}")
  private String username;

  @NotEmpty(message = "{app.login.request.password.require}")
  private String password;
}
