package com.manabie.todotaskapplication.service.ratelimit.impl;

import com.manabie.todotaskapplication.common.constant.TaskActionType;
import com.manabie.todotaskapplication.common.utils.RedisUtils;
import com.manabie.todotaskapplication.service.ratelimit.RateLimitService;
import com.manabie.todotaskapplication.service.userconfig.UserConfigService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.util.Objects;
import java.util.Optional;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
@Service
public class RateLimitServiceImpl implements RateLimitService {
    private static final Logger LOGGER = LoggerFactory.getLogger(RateLimitServiceImpl.class);

    private RedisTemplate<String, Object> redisTemplate;
    private UserConfigService userConfigService;

    @Autowired
    public RateLimitServiceImpl(RedisTemplate<String, Object> redisTemplate, UserConfigService userConfigService) {
        this.redisTemplate = redisTemplate;
        this.userConfigService = userConfigService;
    }

    @Override
    public boolean isValidRateLimit(String userId, LocalDate date, TaskActionType actionType) {
        switch (actionType) {
            case CREATE:
                return isValidRateLimitCreatePerUserPerDay(userId, date);
            default:
                return true;
        }
    }

    private boolean isValidRateLimitCreatePerUserPerDay(String userId, LocalDate date) {
        try {
            Optional<Integer> rateLimit = userConfigService.getRateLimitConfig(userId);
            if (!rateLimit.isPresent()) {
                return false;
            }

            if (rateLimit.get() <= 0) {
                return false;
            }

            Object rs = redisTemplate.opsForHash().get(RedisUtils.getCounterAddedTaskPerUserPerDateKey(userId, date), String.valueOf(userId));
            if (Objects.isNull(rs)) {
                return true;
            }
            Integer counter = (Integer) rs;

            return counter < rateLimit.get();
        } catch (Exception e) {
            LOGGER.error("Error when validate rate limit, userId: {}, date: {}, error: {}", userId, date, e.getMessage());
            throw e;
        }

    }

    @Override
    public boolean increaseCounterAnCheckRateLimit(String userId, LocalDate now) {
        Long counter = redisTemplate.opsForHash().increment(RedisUtils.getCounterAddedTaskPerUserPerDateKey(userId, now), userId, 1);
        Optional<Integer> rateLimit = userConfigService.getRateLimitConfig(userId);
        if (!rateLimit.isPresent() || counter > rateLimit.get()) {
            return false;
        }
        return true;
    }

    @Override
    public long decreaseCounterRateLimit(String userId, LocalDate now) {
        return redisTemplate.opsForHash().increment(RedisUtils.getCounterAddedTaskPerUserPerDateKey(userId, now), userId, -1);
    }
}
