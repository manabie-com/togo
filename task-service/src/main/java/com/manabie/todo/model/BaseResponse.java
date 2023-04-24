package com.manabie.todo.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import lombok.Data;
import org.springframework.http.HttpStatus;

@JsonInclude(JsonInclude.Include.NON_NULL)

@Data
public class BaseResponse<T> {
    private Status status;
    private T data;

    public static <T> BaseResponse<T> ofSucceeded(T data) {
        BaseResponse<T> response = new BaseResponse<>();
        response.data = data;

        Status responseStatus=new Status();
        responseStatus.setCode(HttpStatus.OK.value());
        response.status=responseStatus;
        return response;
    }

    public static <T> BaseResponse<T> ofSucceeded() {
        BaseResponse<T> response = new BaseResponse<>();
        Status responseStatus=new Status();
        responseStatus.setCode(HttpStatus.OK.value());
        response.status=responseStatus;
        return response;
    }

    public static BaseResponse<Void> ofFailed(Integer errorCode) {
        return ofFailed(errorCode, null);
    }

    public static BaseResponse<Void> ofFailed(Integer errorCode, String message) {
        BaseResponse<Void> response = new BaseResponse<>();

        Status responseStatus=new Status();
        responseStatus.setCode(errorCode);
        responseStatus.setMessage(message);
        response.status=responseStatus;
        return response;
    }

    @Data
    public static class Status {
        private String message;
        Integer code;
        boolean success;
    }
}
