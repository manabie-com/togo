package com.manabie.todo.domain;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class User {
  private Long id;
  private String username;
  private String password;
  private Long taskQuote;
  private Boolean isDelete;
}
