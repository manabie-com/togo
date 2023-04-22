package com.manabie.todo.exception;

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
        response.status.code = HttpStatus.OK.value();
        return response;
    }

    public static <T> BaseResponse<T> ofSucceeded() {
        BaseResponse<T> response = new BaseResponse<>();
        response.status.code = HttpStatus.OK.value();
        return response;
    }

    public static BaseResponse<Void> ofFailed(Integer errorCode) {
        return ofFailed(errorCode, null);
    }

    public static BaseResponse<Void> ofFailed(Integer errorCode, String message) {
        BaseResponse<Void> response = new BaseResponse<>();
        response.status.code = errorCode;
        response.status.message = message;
        return response;
    }

    @Data
    public static class Status {
        private String message;
        Integer code;
        boolean success;
    }
}
