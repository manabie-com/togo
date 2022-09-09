package com.manabie.todotaskapplication.common.exception;

import com.manabie.todotaskapplication.common.constant.CustomExceptionCode;
import org.springframework.http.HttpStatus;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public class BusinessException extends AbstractException {
    private static final long serialVersionUID = 1254672957987028725L;

    public BusinessException(CustomExceptionCode code, Throwable cause) {
        super(code, cause);
    }

    public BusinessException(CustomExceptionCode code) {
        super(code);
    }

    @Override
    public HttpStatus getStatus() {
        return HttpStatus.BAD_REQUEST;
    }
}
