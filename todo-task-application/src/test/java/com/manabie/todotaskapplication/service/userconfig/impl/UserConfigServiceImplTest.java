package com.manabie.todotaskapplication.service.userconfig.impl;

import com.github.incu6us.redis.mock.EnableRedisMockTemplate;
import com.manabie.todotaskapplication.common.constant.ConfigType;
import com.manabie.todotaskapplication.common.utils.RedisUtils;
import com.manabie.todotaskapplication.data.model.UserConfig;
import com.manabie.todotaskapplication.repository.userconfig.UserConfigRepository;
import org.junit.Before;
import org.junit.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.junit.runner.RunWith;
import org.mockito.*;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.HashOperations;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.*;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
@RunWith(SpringRunner.class)
@EnableRedisMockTemplate
public class UserConfigServiceImplTest {

    @Spy
    private RedisTemplate<String, Object> redisTemplate;

    @Mock
    private UserConfigRepository userConfigRepository;

    @Mock
    private HashOperations<String, Object, Object> hashOperations;

    @InjectMocks
    UserConfigServiceImpl userConfigService;

    @Test
    public void getUserConfigByType_Success() {
        String userId = "123123";
        ConfigType configType = ConfigType.RATE_LIMIT_ADD_TASK_PER_DAY;
        when(userConfigRepository.findByUserIdAndConfigType(userId, configType.getValue())).thenReturn(new UserConfig(UUID.randomUUID(), userId, configType.getValue(), "5"));
        Optional<UserConfig> userConfig = userConfigService.getUserConfigByType(userId, configType);
        assertTrue(userConfig.isPresent());
        assertEquals("5", userConfig.get().getValue());
    }

    @Test
    public void getUserConfigByType_Failed_Not_Found() {
        String userId = "123123";
        ConfigType configType = ConfigType.RATE_LIMIT_ADD_TASK_PER_DAY;
        when(userConfigRepository.findByUserIdAndConfigType(userId, configType.getValue())).thenReturn(null);
        Optional<UserConfig> userConfig = userConfigService.getUserConfigByType(userId, configType);
        assertFalse(userConfig.isPresent());
    }

    @Test
    public void getRateLimitConfig_Success() {
        String userId = "123123";
        ConfigType configType = ConfigType.RATE_LIMIT_ADD_TASK_PER_DAY;
        when(redisTemplate.opsForHash()).thenReturn(hashOperations);
        when(hashOperations.get(RedisUtils.getRateLimitConfigKey(userId), userId)).thenReturn(null);
        when(userConfigRepository.findByUserIdAndConfigType(userId, configType.getValue())).thenReturn(new UserConfig(UUID.randomUUID(), userId, configType.getValue(), "5"));

        Optional<Integer> rateLimitConfig = userConfigService.getRateLimitConfig(userId);
        assertTrue(rateLimitConfig.isPresent());
        assertEquals(5, rateLimitConfig.get());
    }

    @Test
    public void getRateLimitConfig_Success_With_Redis_Value() {
        String userId = "123123";
        ConfigType configType = ConfigType.RATE_LIMIT_ADD_TASK_PER_DAY;
        int rateLimitValueRedis = 10;
        when(redisTemplate.opsForHash()).thenReturn(hashOperations);
        when(hashOperations.get(RedisUtils.getRateLimitConfigKey(userId), userId)).thenReturn(rateLimitValueRedis);
        Optional<Integer> rateLimitConfig = userConfigService.getRateLimitConfig(userId);
        assertTrue(rateLimitConfig.isPresent());
        assertEquals(rateLimitValueRedis, rateLimitConfig.get());
    }

    @Test
    public void getRateLimitConfig_Success_With_Not_Config_Yet() {
        String userId = "123123";
        ConfigType configType = ConfigType.RATE_LIMIT_ADD_TASK_PER_DAY;
        when(userConfigRepository.findByUserIdAndConfigType(userId, configType.getValue())).thenReturn(null);
        Optional<Integer> rateLimitConfig = userConfigService.getRateLimitConfig(userId);
        assertTrue(rateLimitConfig.isPresent());
        assertEquals(-1, rateLimitConfig.get());
    }

    @Test
    public void getRateLimitConfig_Success_With_Failed_To_Get_From_Redis() {

    }

}