package com.manabie.todo.model;

import jakarta.validation.constraints.NotBlank;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CreateTaskRequest {
    @NotBlank(message = "title is mandatory")
    private String title;

    private String description;

    @NotBlank(message = "owner is mandatory")
    private Long owner;
}
