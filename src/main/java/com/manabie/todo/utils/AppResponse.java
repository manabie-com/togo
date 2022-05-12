package com.manabie.todo.utils;

import lombok.Builder;
import lombok.Data;

import java.util.List;

@Data
@Builder
public class AppResponse<T> {

  private AppCode code;
  private String message;
  private T data;

  public static <T> AppResponse<T> ok(T data) {
    return AppResponse.<T>builder()
        .code(AppCode.SUCCESS)
        .message(AppCode.SUCCESS.getMessage())
        .data(data)
        .build();
  }

  public static AppResponse<List<String>> userError(List<String> data) {
    return AppResponse.<List<String>>builder()
        .code(AppCode.USER_ERROR)
        .message(AppCode.USER_ERROR.getMessage())
        .data(data)
        .build();
  }
}
