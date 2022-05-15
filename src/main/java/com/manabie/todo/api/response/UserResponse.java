package com.manabie.todo.api.response;

import lombok.Builder;
import lombok.Data;

@Builder
@Data
public class UserResponse {
  private long id;
  private String username;
  private long taskQuote;
}
