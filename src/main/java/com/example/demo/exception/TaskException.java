package com.example.demo.exception;

import lombok.Data;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@Data
public class TaskException extends RuntimeException {
    private static final long serialVersionUID = 1L;
    private String message;
    public TaskException(String message) {
        this.message = message;
    }
}
