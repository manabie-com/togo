package com.todo.core.commons.model;

public class GenericResponse<T> {
    public T data;
    public String message;

    public GenericResponse(T data, String message) {
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

