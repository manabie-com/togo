package com.manabie.todotaskapplication.common.exception;

import com.manabie.todotaskapplication.common.constant.CustomExceptionCode;
import org.springframework.http.HttpStatus;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public class ValidationException extends AbstractException {
    public ValidationException(CustomExceptionCode code, Throwable cause) {
        super(code, cause);
    }

    public ValidationException(CustomExceptionCode code) {
        super(code);
    }

    @Override
    public HttpStatus getStatus() {
        switch (this.getCode()) {
            case RATE_LIMIT_EXCEPTION:
                return HttpStatus.TOO_MANY_REQUESTS;
            default:
                return HttpStatus.BAD_REQUEST;
        }
    }
}
