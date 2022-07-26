package com.example.demo.controller;

import com.example.demo.exception.TaskException;
import com.example.demo.exception.TaskExceptionResponse;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

@ControllerAdvice
@Slf4j
public class TogoExceptionController extends ResponseEntityExceptionHandler {

    @ExceptionHandler(TaskException.class)
    public ResponseEntity taskExceptionHandler(TaskException exception) {
        log.error(exception.getMessage());
        return ResponseEntity
                .status(HttpStatus.BAD_REQUEST)
                .body(new TaskExceptionResponse(exception.getMessage()));
    }
}
