package com.manabie.togo.dto;

import lombok.Builder;
import lombok.Data;

import javax.validation.constraints.NotNull;

@Data
@Builder
public class ToDoRequest {

    @NotNull
    private String userId;

    @NotNull
    private String title;

    @NotNull
    private String description;

    @NotNull
    private String toDoDate;
}
