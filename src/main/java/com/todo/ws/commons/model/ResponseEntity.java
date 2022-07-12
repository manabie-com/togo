package com.todo.ws.commons.model;

public class ResponseEntity<T> {
    public T data;
    public String message;

    public ResponseEntity(T data, String message) {
        this.data = data;
        this.message = message;
    }

    public T getData() {
        return data;
    }

    public String getMessage() {
        return message;
    }
}

