package com.example.demo.model;

import lombok.Data;

@Data
public class RegisterRequest {
    private String username;
    private String password;
    private Long limit;
}
