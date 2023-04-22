package com.manabie.todo.service;

public interface TaskLimitService {
    boolean checkLimitAndIncreaseCounter(Long userId, Integer maxTaxPerDay);
}
