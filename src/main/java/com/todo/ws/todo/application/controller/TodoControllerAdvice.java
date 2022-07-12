package com.todo.ws.todo.application.controller;

import com.todo.ws.commons.enums.TodoEnum;
import com.todo.ws.commons.model.ResponseEntity;
import com.todo.ws.todo.exception.LimitForTodayReachedException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice
public class TodoControllerAdvice {

    @ExceptionHandler(LimitForTodayReachedException.class)
    public ResponseEntity<String> handleLimitReached(LimitForTodayReachedException ex) {
        return new ResponseEntity<>(ex.getMessage(), TodoEnum.LIMIT_REACHED.getContent());
    }
}
