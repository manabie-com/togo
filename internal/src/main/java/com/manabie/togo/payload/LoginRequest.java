package com.manabie.togo.payload;

import javax.validation.constraints.NotBlank;

import lombok.Data;

@Data
public class LoginRequest {

    @NotBlank
    private String user_id;

    @NotBlank
    private String password;
}
