package com.manabie.todo.exception;

import com.manabie.todo.model.BaseResponse;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice
public class ManabieExceptionHandler {
    @ExceptionHandler({ManabieException.class})
    public ResponseEntity<BaseResponse> vuiException(ManabieException ex) {
        BaseResponse body = BaseResponse.ofFailed(ex.getCode(), ex.getMessage());
        return ResponseEntity.status(ex.getHttpStatus()).body(body);
    }

    @ExceptionHandler({Exception.class})
    @ResponseStatus(value = HttpStatus.INTERNAL_SERVER_ERROR)
    public ResponseEntity<BaseResponse> exception(Exception ex) {
        BaseResponse body = BaseResponse.ofFailed(HttpStatus.INTERNAL_SERVER_ERROR.value(), ex.getMessage());
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(body);
    }
}
