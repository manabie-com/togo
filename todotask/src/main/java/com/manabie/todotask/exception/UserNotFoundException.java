package com.manabie.todotask.exception;

import lombok.Data;

import java.util.Date;

@Data
public class UserNotFoundException extends RuntimeException{
    private Integer userId;
    public UserNotFoundException(Integer userId){
        super("User with id " + userId + " does not exist");
        this.userId = userId;
    }
}
