package com.manabie.todotask.exception;

import lombok.Data;

import java.time.ZonedDateTime;

@Data
public class ReachLimitException extends RuntimeException{
    protected Integer userId;
    protected ZonedDateTime dateTime;
    public ReachLimitException(Integer userId, ZonedDateTime date){
        super("User " + userId + " has reached max task limit at date " + date);
        this.userId = userId;
        this.dateTime = date;
    }
}
