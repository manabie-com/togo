package com.todo.core.commons;

public enum Messages {

    SAVE_SUCCESSFUL ("Save Successful"),
    USER_CREATE_SUCCESSFUL ("User Create Successful"),
    ERROR_ON_LOGIN ("Error on Login"),
    PAGES_RETRIEVE_SUCCESSFUL ("Pages Retrieved for User"),
    LIMIT_REACHED ("Limit for user reached"),
    USERNAME_TAKEN ("Username is already in use");

    private final String content;

    Messages(String content) {
        this.content = content;
    }

    public String getContent() {
        return content;
    }

}
