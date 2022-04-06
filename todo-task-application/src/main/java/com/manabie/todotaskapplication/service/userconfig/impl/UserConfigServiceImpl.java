package com.manabie.todotaskapplication.service.userconfig.impl;

import com.manabie.todotaskapplication.common.constant.ConfigType;
import com.manabie.todotaskapplication.common.utils.RedisUtils;
import com.manabie.todotaskapplication.data.model.UserConfig;
import com.manabie.todotaskapplication.repository.userconfig.UserConfigRepository;
import com.manabie.todotaskapplication.service.userconfig.UserConfigService;
import org.apache.commons.lang3.StringUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

import java.util.Objects;
import java.util.Optional;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
@Service
public class UserConfigServiceImpl implements UserConfigService {
    private static final Logger LOGGER = LoggerFactory.getLogger(UserConfigServiceImpl.class);
    private UserConfigRepository userConfigRepository;
    private RedisTemplate<String, Object> redisTemplate;

    @Autowired
    public UserConfigServiceImpl(UserConfigRepository userConfigRepository, RedisTemplate<String, Object> redisTemplate) {
        this.userConfigRepository = userConfigRepository;
        this.redisTemplate = redisTemplate;
    }

    @Override
    public Optional<UserConfig> getUserConfigByType(String userId, ConfigType configType) {
        return Optional.ofNullable(userConfigRepository.findByUserIdAndConfigType(userId, configType.getValue()));
    }

    @Override
    public Optional<Integer> getRateLimitConfig(String userId) {
        try {
            Object rateLimitValue = redisTemplate.opsForHash().get(RedisUtils.getRateLimitConfigKey(userId), userId);
            if (Objects.nonNull(rateLimitValue)) {
                String rateLimit = String.valueOf(rateLimitValue);
                if (StringUtils.isNumeric(rateLimit)) {
                    return Optional.of(Integer.valueOf(rateLimit));
                }
            }
        } catch (Exception e) {
            LOGGER.warn("Failed to get rate limit from redis, error: {}", e.getMessage());
        }


        Optional<UserConfig> userConfig = getUserConfigByType(userId, ConfigType.RATE_LIMIT_ADD_TASK_PER_DAY);
        if (!userConfig.isPresent() || !StringUtils.isNumeric(userConfig.get().getValue())) {
            return Optional.of(-1);
        }

        int rateLimitMax = Integer.parseInt(userConfig.get().getValue());

        redisTemplate.opsForHash().put(RedisUtils.getRateLimitConfigKey(userId), userId, rateLimitMax);

        return Optional.of(rateLimitMax);
    }
}
