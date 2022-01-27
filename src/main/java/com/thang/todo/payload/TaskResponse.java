package com.thang.todo.payload;

public class TaskResponse {
    private String accessToken;
    private String tokenType = "Bearer";

    public TaskResponse(String accessToken) {
        this.accessToken = accessToken;
    }

    public String getAccessToken() {
        return this.accessToken;
    }

    public void setAccessToken(String accessToken) {
        this.accessToken = accessToken;
    }

    public String getTokenType() {
        return this.tokenType;
    }
}
