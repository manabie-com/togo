package com.todo.ws.core.auth.dto;

import org.springframework.security.crypto.password.PasswordEncoder;

public class RegRequestDto {

    private String username;
    private String password;
    private int todoLimit;

    public RegRequestDto(String username, String password, int todoLimit) {
        this.username = username;
        this.password = password;
        this.todoLimit = todoLimit;
    }

    public String getUsername() {
        return username;
    }

    public String getPassword() {
        return password;
    }

    public void doEncodePassword(PasswordEncoder passwordEncoder) {
        this.password = passwordEncoder.encode(password);
    }

    public int getTodoLimit() {
        return todoLimit;
    }
}
