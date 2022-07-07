package com.todo.ws.common.enums;

public enum AuthMessageEnum {
	ERROR_ON_LOGIN ("Login failed"),
	USERNAME_EXISTS ("Username already exists"),
	USER_CREATE_SUCCESSFUL ("User Registration Successful");
	
    private final String content;

	AuthMessageEnum(String content) {
        this.content = content;
    }

    public String getContent() {
        return content;
    }
}
