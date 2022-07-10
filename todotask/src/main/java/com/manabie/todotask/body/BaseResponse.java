package com.manabie.todotask.body;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
public class BaseResponse {
    private int code;
    private String message;
    private Object data;

    public BaseResponse(int code, String message){
        this.code = code;
        this.message = message;
    }

    public BaseResponse withData(Object data){
        this.data = data;
        return this;
    }
}
