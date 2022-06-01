package com.example.demo.exception;

import lombok.Data;

@Data
public class TaskException extends RuntimeException {
    private static final long serialVersionUID = 1L;
    private String message;
    public TaskException(String message) {
        this.message = message;
    }
}
