package com.manabie.todo.config;

import com.manabie.todo.constant.CacheKey;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.integration.redis.util.RedisLockRegistry;

@Configuration
public class RedisConfig {
    @Bean(destroyMethod = "destroy")
    public RedisLockRegistry userLockRegistry(RedisConnectionFactory redisConnectionFactory) {
        return new RedisLockRegistry(redisConnectionFactory, CacheKey.USER_LOCK);
    }
}
