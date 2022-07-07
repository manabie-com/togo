package com.todo.ws.model;

public class ResponseEntity<T> {
	private T data;
    private String message;

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
    
    public void setData(T data) {
    	this.data = data;
    }
    
    public void setMessage(String message) {
    	this.message = message;
    }
}
