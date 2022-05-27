package com.todo.core.user.application.dto;

import org.springframework.security.crypto.password.PasswordEncoder;

public class UserRegistrationDTO {

    private String username;
    private String password;
    private int todoLimit;

    public UserRegistrationDTO(String username, String password, int todoLimit) {
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
