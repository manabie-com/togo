package com.manabie.todo.domain;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class Task {
  private Long id;
  private Long userId;
  private String title;
  private String description;
  private Timestamp datetimeCreated;
  private Timestamp datetimeEdited;
  private Boolean isDelete;
}
