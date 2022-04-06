package com.manabie.todotaskapplication.common.exception;

import com.manabie.todotaskapplication.common.constant.CustomExceptionCode;
import org.springframework.http.HttpStatus;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public class TechnicalException extends AbstractException {
    private static final long serialVersionUID = 1254672957987028725L;

    public TechnicalException(CustomExceptionCode code, Throwable cause) {
        super(code, cause);
    }

    public TechnicalException(CustomExceptionCode code) {
        super(code);
    }

    @Override
    public HttpStatus getStatus() {
        return HttpStatus.INTERNAL_SERVER_ERROR;
    }
}
