package com.todo.ws.commons.enums;

public enum AuthenticationEnum {

    USER_CREATE_SUCCESSFUL ("User Create Successful"),
    ERROR_ON_LOGIN ("Error on Login"),
    USERNAME_TAKEN ("Username is already in use");
	

    private final String content;

    AuthenticationEnum(String content) {
        this.content = content;
    }

    public String getContent() {
        return content;
    }
}
