package com.example.demo.model;

import lombok.Data;

@Data
public class CompleteTaskRequest {
    private long id;
    private Boolean isTaskCompleted;
}
