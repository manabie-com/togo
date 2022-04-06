
package com.manabie.todotaskapplication.common.exception.hanlder;

import com.manabie.todotaskapplication.common.exception.AbstractException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.MessageSource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.util.CollectionUtils;

import java.util.*;
import java.util.stream.Collectors;

public abstract class AbstractExceptionHandler {
    private static final Logger logger = LoggerFactory.getLogger(AbstractExceptionHandler.class);

    /**
     * Handle the Exception in RestController during implementation
     *
     * @param exception
     * @param locale
     * @return
     * @throws Throwable
     */
    public ResponseEntity<Object> handle(AbstractException exception, Locale locale) {
        if (logger.isDebugEnabled()) {
            logger.info(exception.getMessage(), exception);
        }

        HttpStatus httpStatus = exception.getStatus();
        HttpHeaders headers = new HttpHeaders();
        headers.add(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE);
        return new ResponseEntity<>(exception.getCause().getMessage(), headers, httpStatus);
    }

    protected abstract MessageSource getMessageSource();
}
