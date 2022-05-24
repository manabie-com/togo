package com.todo.core.todo.exception;

public class LimitForTodayReachedException extends RuntimeException {

    public LimitForTodayReachedException(String message) {
        super(message);
    }

    public LimitForTodayReachedException(String message, Throwable cause) {
        super(message, cause);
    }
}
