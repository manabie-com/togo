package com.manabie.togo.payload;

import lombok.Data;

@Data
public class LoginResponse {

    private String data;

    public LoginResponse(String data) {
        this.data = data;
    }
}
