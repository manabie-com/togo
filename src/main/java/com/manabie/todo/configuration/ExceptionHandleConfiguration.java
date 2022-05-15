package com.manabie.todo.configuration;

import com.manabie.todo.exception.AppException;
import com.manabie.todo.exception.TaskLimitException;
import com.manabie.todo.exception.UserUnauthorizedException;
import com.manabie.todo.utils.AppCode;
import com.manabie.todo.utils.AppResponse;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.context.MessageSource;
import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.ObjectError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.Collections;
import java.util.List;
import java.util.Locale;
import java.util.Objects;

@AllArgsConstructor
@ControllerAdvice
@Slf4j
public class ExceptionHandleConfiguration extends ResponseEntityExceptionHandler {

  private final MessageSource messageSource;

  @Override
  protected ResponseEntity<Object> handleMethodArgumentNotValid(MethodArgumentNotValidException ex,
                                                                HttpHeaders headers, HttpStatus status, WebRequest request) {
    return ResponseEntity.badRequest().body(AppResponse
        .userError(
            ex.getAllErrors().stream()
                .map(ObjectError::getDefaultMessage).filter(Objects::nonNull).toList()
        ));
  }

  @ExceptionHandler({UserUnauthorizedException.class})
  public ResponseEntity<Object> handleUserUnauthorizedException(
      UserUnauthorizedException appException) {
    return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body(AppResponse
        .userError(
            bindingMessage(appException.getCode())
        ));
  }

  @ExceptionHandler({AppException.class})
  public ResponseEntity<Object> handleAppException(AppException appException) {
    return ResponseEntity.badRequest().body(AppResponse
        .userError(
            bindingMessage(appException.getCode())
        ));
  }

  @ExceptionHandler({TaskLimitException.class})
  public ResponseEntity<Object> handleTaskLimitException(TaskLimitException taskLimitException) {
    return ResponseEntity.badRequest().body(AppResponse
        .userError(
            bindingMessage(taskLimitException.getCode())
        ));
  }

  @ExceptionHandler(EmptyResultDataAccessException.class)
  public ResponseEntity<Object> handleEmptyResultDataAccessException(EmptyResultDataAccessException ex){
    log.debug(ex.getMessage());
    return ResponseEntity.notFound().build();
  }

  @ExceptionHandler(Exception.class)
  public ResponseEntity<Object> handleCommonException(Exception ex,
                                                      HttpServletRequest request, HttpServletResponse response) {
    if (ex instanceof NullPointerException) {
      ex.printStackTrace();
      return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
    }
    log.debug(ex.getMessage());
    return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).build();
  }

  private List<String> bindingMessage(AppCode appException) {
    log.debug(messageSource.getMessage(appException.getMessage(), null, Locale.getDefault()));
    return Collections.singletonList(
        messageSource.getMessage(appException.getMessage(), null,
            Locale.getDefault()));
  }
}
