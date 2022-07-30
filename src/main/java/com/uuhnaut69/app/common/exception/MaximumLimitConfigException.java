package com.uuhnaut69.app.common.exception;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

/**
 * @author uuhnaut
 */
@ResponseStatus(HttpStatus.BAD_REQUEST)
public class MaximumLimitConfigException extends RuntimeException {

  public MaximumLimitConfigException(String message) {
    super(message);
  }
}
