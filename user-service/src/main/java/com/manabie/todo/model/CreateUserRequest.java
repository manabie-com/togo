package com.manabie.todo.model;

import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CreateUserRequest {
    @NotBlank(message = "username is mandatory")
    private String username;

    @Min(value=1, message="maxTaskPerDay must be equal or greater than 1")
    @Max(value=45, message="maxTaskPerDay must be equal or less than 10")
    private Integer maxTaskPerDay;
}
