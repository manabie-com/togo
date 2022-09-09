package com.manabie.todotaskapplication.service.ratelimit;

import com.manabie.todotaskapplication.common.constant.TaskActionType;
import com.manabie.todotaskapplication.common.utils.RedisUtils;

import java.time.LocalDate;
import java.util.Objects;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
public interface RateLimitService {
    boolean isValidRateLimit(String userId, LocalDate date, TaskActionType actionType);

    boolean increaseCounterAnCheckRateLimit(String userId, LocalDate now);

    long decreaseCounterRateLimit(String userId, LocalDate now);
}
