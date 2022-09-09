package com.manabie.todotaskapplication.config.property;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;
import org.springframework.util.SocketUtils;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
@Configuration
public class RedisProperties {
    public static final int redisPort = SocketUtils.findAvailableTcpPort();
}
