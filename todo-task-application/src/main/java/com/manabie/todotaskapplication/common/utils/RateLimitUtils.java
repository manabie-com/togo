package com.manabie.todotaskapplication.common.utils;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public class RateLimitUtils {
    private RateLimitUtils() {
    }

    private static final int MAX_ITEM_PER_HASH_BUCKET = 1000;
    public static final String RATE_LIMIT_KEY_FORMATTER = "todo_task_%s_%s";//todo_task_{bucketid}_{date}



    public static int getHashBucketId(String userId) {
        return Integer.parseInt(userId) / MAX_ITEM_PER_HASH_BUCKET;
    }
}
