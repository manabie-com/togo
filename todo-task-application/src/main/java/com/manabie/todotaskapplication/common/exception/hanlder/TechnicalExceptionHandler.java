package com.manabie.todotaskapplication.common.exception.hanlder;

import com.manabie.todotaskapplication.common.exception.AbstractException;
import com.manabie.todotaskapplication.common.exception.TechnicalException;
import com.manabie.todotaskapplication.common.exception.ValidationException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.MessageSource;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.Locale;

@ControllerAdvice
@Order(value = Ordered.HIGHEST_PRECEDENCE)
public class TechnicalExceptionHandler extends AbstractExceptionHandler {

    private final MessageSource validationMessages;

    @Autowired
    public TechnicalExceptionHandler(MessageSource validationMessages) {
        this.validationMessages = validationMessages;
    }

    @ExceptionHandler(TechnicalException.class)
    @ResponseBody
    public ResponseEntity<Object> handleValidationException(AbstractException ex, Locale locale) {
        return super.handle(ex, locale);
    }

    @Override
    protected MessageSource getMessageSource() {
        return validationMessages;
    }
}
