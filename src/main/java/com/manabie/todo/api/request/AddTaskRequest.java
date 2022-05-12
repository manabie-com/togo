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
public class AddTaskRequest {
  @NotEmpty(message = "{app.task.request.title.require}")
  private String title;

  @NotEmpty(message = "{app.task.request.description.require}")
  private String description;
}
