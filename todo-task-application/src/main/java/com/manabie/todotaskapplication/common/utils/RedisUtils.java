package com.manabie.todotaskapplication.common.utils;

import org.springframework.util.SocketUtils;

import java.time.LocalDate;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public class RedisUtils {
    public static final String COUNTER_ADDED_TASK_PER_USER_KEY_FORMATTER = "counter_added_task_%s_%s";//todo_task_{bucketid}_{date}
    public static final String RATE_LIMIT_CONFIG_KEY_FORMATTER = "user_config_rate_limit_%s";//user_config_rate_limit_{bucketid}

    public static String getCounterAddedTaskPerUserPerDateKey(String userId, LocalDate date) {
        return String.format(COUNTER_ADDED_TASK_PER_USER_KEY_FORMATTER, RateLimitUtils.getHashBucketId(userId), date);
    }

    public static String getRateLimitConfigKey(String userId) {
        return String.format(RATE_LIMIT_CONFIG_KEY_FORMATTER, RateLimitUtils.getHashBucketId(userId));
    }
}
