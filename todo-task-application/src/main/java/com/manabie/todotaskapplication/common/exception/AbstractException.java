package com.manabie.todotaskapplication.common.exception;

import com.manabie.todotaskapplication.common.constant.CustomExceptionCode;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;
import org.springframework.http.HttpStatus;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
@Getter
public abstract class AbstractException extends RuntimeException {
    private static final long serialVersionUID = 1254672957987028725L;
    private CustomExceptionCode code;

    public AbstractException(CustomExceptionCode code, Throwable cause) {
        super(cause);
        this.code = code;
    }
    public AbstractException(CustomExceptionCode code) {
        this.code = code;
    }
    public abstract HttpStatus getStatus();
}
