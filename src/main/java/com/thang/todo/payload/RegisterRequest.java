package com.thang.todo.payload;

import javax.validation.constraints.NotBlank;

public class RegisterRequest {
    @NotBlank
    private String username;

    @NotBlank
    private String password;

    public String getUsername() {
        return this.username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return this.password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
}

