package com.manabie.todo.model;

import com.manabie.todo.constant.TaskStatus;
import lombok.Getter;
import lombok.Setter;

import java.time.LocalDateTime;

@Getter
@Setter
public class TaskModel {
    private Long id;
    private String title;
    private String description;
    private Long owner;
    private TaskStatus status;
    private LocalDateTime createdAt;
    private String createdBy;
    private LocalDateTime updatedAt;
    private String updatedBy;
}
