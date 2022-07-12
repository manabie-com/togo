package com.todo.ws.user.application.controller;

import com.todo.ws.commons.enums.AuthenticationEnum;
import com.todo.ws.commons.model.ResponseEntity;
import com.todo.ws.user.exception.UserAlreadyExistsException;
import com.todo.ws.user.exception.UserDoesNotExistException;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice
public class UserControllerAdvice {

    @ExceptionHandler(UserDoesNotExistException.class)
    @ResponseStatus(HttpStatus.OK)
    public ResponseEntity<String> handleBadUser(UserDoesNotExistException ex) {
        return new ResponseEntity<>(ex.getMessage(), AuthenticationEnum.ERROR_ON_LOGIN.getContent());
    }

    @ExceptionHandler(UserAlreadyExistsException.class)
    @ResponseStatus(HttpStatus.OK)
    public ResponseEntity<String> handleTakenUser(UserAlreadyExistsException ex) {
        return new ResponseEntity<>(ex.getMessage(), AuthenticationEnum.USERNAME_TAKEN.getContent());
    }
}
