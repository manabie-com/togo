package com.manabie.todo.exception;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import org.springframework.http.HttpStatus;

@Builder
@Data
@AllArgsConstructor
public class ManabieException extends RuntimeException {
    private int code;
    private String message;
    private HttpStatus httpStatus;
}
