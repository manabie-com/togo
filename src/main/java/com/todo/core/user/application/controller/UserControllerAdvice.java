package com.todo.core.user.application.controller;

import com.todo.core.commons.Messages;
import com.todo.core.commons.model.GenericResponse;
import com.todo.core.user.exception.UserAlreadyExistsException;
import com.todo.core.user.exception.UserDoesNotExistException;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice
public class UserControllerAdvice {

    @ExceptionHandler(UserDoesNotExistException.class)
    @ResponseStatus(HttpStatus.OK)
    public GenericResponse<String> handleBadUser(UserDoesNotExistException ex) {
        return new GenericResponse<>(ex.getMessage(), Messages.ERROR_ON_LOGIN.getContent());
    }

    @ExceptionHandler(UserAlreadyExistsException.class)
    @ResponseStatus(HttpStatus.OK)
    public GenericResponse<String> handleTakenUser(UserAlreadyExistsException ex) {
        return new GenericResponse<>(ex.getMessage(), Messages.USERNAME_TAKEN.getContent());
    }
}
