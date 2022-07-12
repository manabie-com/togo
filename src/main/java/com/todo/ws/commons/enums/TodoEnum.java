package com.todo.ws.commons.enums;

public enum TodoEnum {
    SAVE_SUCCESSFUL ("Save Successful"),
    PAGES_RETRIEVE_SUCCESSFUL ("Pages Retrieved for User"),
    LIMIT_REACHED ("Limit for user reached");

    private final String content;

    TodoEnum(String content) {
        this.content = content;
    }

    public String getContent() {
        return content;
    }
}
