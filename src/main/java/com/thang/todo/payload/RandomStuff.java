package com.thang.todo.payload;

public class RandomStuff {
    private String message;

    public void setMessage(String message) {
        this.message = message;
    }

    public String getMessage() {
        return this.message;
    }

    public RandomStuff(String message) {
        this.message = message;
    }
}