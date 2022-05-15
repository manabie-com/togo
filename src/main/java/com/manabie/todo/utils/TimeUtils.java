package com.manabie.todo.utils;

import java.sql.Timestamp;
import java.util.Calendar;

import static com.manabie.todo.utils.Constants.MILLISECONDS_OF_DAY;

public class TimeUtils {
  public static boolean isCreatedToday(Timestamp timestamp){
    //check diff < 24h
    if(System.currentTimeMillis() - timestamp.getTime() >= MILLISECONDS_OF_DAY)
      return false;

    Calendar calendar = Calendar.getInstance();
    calendar.setTime(timestamp);
    int dayOfMonth = calendar.get(Calendar.DAY_OF_MONTH);

    Calendar currTime = Calendar.getInstance();
    currTime.setTime(new Timestamp(System.currentTimeMillis()));
    int currDayOfMonth = currTime.get(Calendar.DAY_OF_MONTH);

    return dayOfMonth == currDayOfMonth;
  }
}
