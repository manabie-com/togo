package com.uuhnaut69.app.todo.model.dto;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;
import javax.validation.constraints.Size;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * @author uuhnaut
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class TodoRequest {

  @Size(max = 255, message = "Maximum task length is 255")
  @NotBlank(message = "Task can not blank")
  private String task;

  @NotNull(message = "User id can not null")
  private Long userId;
}
