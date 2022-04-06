package com.manabie.todotaskapplication.config;

import com.manabie.todotaskapplication.config.property.RedisProperties;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.annotation.Configuration;
import org.springframework.stereotype.Component;
import redis.embedded.RedisServer;

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;
import java.io.IOException;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
@Configuration
public class EmbeddedRedis {
    private static final Logger LOGGER = LoggerFactory.getLogger(EmbeddedRedis.class);

    private RedisServer redisServer;

    @PostConstruct
    public void startRedis() throws IOException {
        redisServer = new RedisServer(RedisProperties.redisPort);
        redisServer.start();
        LOGGER.info("EmbeddedRedis start at, port: {}", RedisProperties.redisPort);
    }

    @PreDestroy
    public void stopRedis() {
        redisServer.stop();
    }
}
