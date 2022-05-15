package com.manabie.todo.utils;

import org.junit.jupiter.api.Test;

import java.sql.Time;
import java.sql.Timestamp;
import java.util.Calendar;
import java.util.Date;

import static org.junit.jupiter.api.Assertions.*;

class TimeUtilsTest {

  @Test
  void given1DayDiff_whenCheck_thenFalse() {
    //given
    Calendar calendar = Calendar.getInstance();
    calendar.setTime(new Date());
    calendar.set(Calendar.DAY_OF_MONTH,-1);

    //when then
    assertFalse(TimeUtils.isCreatedToday(new Timestamp(calendar.getTime().getTime())));
  }

  @Test
  void givenCurrentDay_whenCheck_thenTrue() {
    //given
    Calendar calendar = Calendar.getInstance();
    calendar.setTime(new Date());

    //when then
    assertTrue(TimeUtils.isCreatedToday(new Timestamp(calendar.getTime().getTime())));
  }
}