package com.todo.model;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
public class TodoTaskDTO {
    private String content;
    private boolean isComplete;
}
