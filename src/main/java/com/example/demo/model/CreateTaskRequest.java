package com.example.demo.model;

import lombok.Data;

@Data
public class CreateTaskRequest {
    private String taskDetail;
    private Boolean isCompleted;
}
