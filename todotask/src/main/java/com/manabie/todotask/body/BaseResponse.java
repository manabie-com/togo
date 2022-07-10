package com.manabie.todotask.body;

import lombok.Data;

@Data
public class BaseResponse {
    private int status;
    private String message;
    private Object data;

    public BaseResponse(int status, String message) {
        this.status = status;
        this.message = message;
    }

    public BaseResponse withData(Object data) {
        this.data = data;
        return this;
    }
}
