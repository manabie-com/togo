package com.todo.core.todo.application.controller;

import com.todo.core.commons.Messages;
import com.todo.core.commons.model.GenericResponse;
import com.todo.core.todo.exception.LimitForTodayReachedException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice
public class TodoControllerAdvice {

    @ExceptionHandler(LimitForTodayReachedException.class)
    public GenericResponse<String> handleLimitReached(LimitForTodayReachedException ex) {
        return new GenericResponse<>(ex.getMessage(), Messages.LIMIT_REACHED.getContent());
    }
}
