package com.manabie.todo.utils;

public class Constants {

  public static final String PASSWORD_SECURITY_PATTERN = "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{8,}$";

  public static final int JWT_TOKEN_INDEX = 1;
  public static final String PREFIX_TOKEN = "Bearer ";
  public static final String USER_ID_CLAIM = "userId";
  public static final String TASK_QUOTE_CLAIM = "taskQuote";
  public static final String USERNAME_CLAIM = "username";

  public static final long MILLISECONDS_OF_DAY = 24*60*60*1000;

}
