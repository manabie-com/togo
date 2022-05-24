package com.todo.core.commons;

public enum Messages {

    SAVE_SUCCESSFUL ("Save Successful"),
    USER_CREATE_SUCCESSFUL ("User Create Successful");

    private final String content;

    Messages(String content) {
        this.content = content;
    }

    public String getContent() {
        return content;
    }

}
