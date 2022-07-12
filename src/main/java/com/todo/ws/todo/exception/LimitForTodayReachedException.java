package com.todo.ws.todo.exception;

public class LimitForTodayReachedException extends RuntimeException {

    public LimitForTodayReachedException(String message) {
        super(message);
    }

    public LimitForTodayReachedException(String message, Throwable cause) {
        super(message, cause);
    }
}
