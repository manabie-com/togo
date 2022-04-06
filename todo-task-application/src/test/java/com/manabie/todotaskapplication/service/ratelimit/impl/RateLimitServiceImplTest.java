package com.manabie.todotaskapplication.service.ratelimit.impl;

import com.manabie.todotaskapplication.common.constant.TaskActionType;
import com.manabie.todotaskapplication.common.utils.RedisUtils;
import com.manabie.todotaskapplication.service.userconfig.impl.UserConfigServiceImpl;
import org.junit.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.Spy;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.data.redis.core.HashOperations;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.test.context.junit4.SpringRunner;

import java.time.LocalDate;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.when;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
@RunWith(SpringRunner.class)
public class RateLimitServiceImplTest {
    @Spy
    private RedisTemplate<String, Object> redisTemplate;
    @Mock
    private HashOperations<String, Object, Object> hashOperations;
    @InjectMocks
    RateLimitServiceImpl rateLimitService;
    @Mock
    private UserConfigServiceImpl userConfigService;

    @Test
    public void isValidRateLimit_Success_With_Other_Action_With_Task() {
        boolean isValid = rateLimitService.isValidRateLimit("12312", LocalDate.now(), TaskActionType.UPDATE);
        assertTrue(isValid);
    }

    @Test
    public void isValidRateLimit_False_Action_Create_Task_Not_User_Config_Yet() {
        String userId = "123123";
        when(userConfigService.getRateLimitConfig(userId)).thenReturn(Optional.empty());
        boolean isValid = rateLimitService.isValidRateLimit(userId, LocalDate.now(), TaskActionType.CREATE);
        assertFalse(isValid);
    }

    @Test
    public void isValidRateLimit_False_Action_Create_Task_Not_User_Config_Negative() {
        String userId = "123123";
        when(userConfigService.getRateLimitConfig(userId)).thenReturn(Optional.of(-1));
        boolean isValid = rateLimitService.isValidRateLimit(userId, LocalDate.now(), TaskActionType.CREATE);
        assertFalse(isValid);
    }

    @Test
    public void isValidRateLimit_True_Action_Create_Task_Redis_Null_Counter() {
        String userId = "123123";
        Optional<Integer> rateLimitConfig = Optional.of(10);
        when(userConfigService.getRateLimitConfig(userId)).thenReturn(rateLimitConfig);
        when(redisTemplate.opsForHash()).thenReturn(hashOperations);
        when(hashOperations.get(anyString(), anyString())).thenReturn(null);
        boolean isValid = rateLimitService.isValidRateLimit(userId, LocalDate.now(), TaskActionType.CREATE);
        assertTrue(isValid);
    }

    @Test
    public void isValidRateLimit_True_Action_Create_Task_Redis_Counter_Smaller_Than_Rate_Limit_Config() {
        String userId = "123123";
        Optional<Integer> rateLimitConfig = Optional.of(10);
        when(userConfigService.getRateLimitConfig(userId)).thenReturn(rateLimitConfig);
        when(redisTemplate.opsForHash()).thenReturn(hashOperations);
        when(hashOperations.get(anyString(), anyString())).thenReturn(5);
        boolean isValid = rateLimitService.isValidRateLimit(userId, LocalDate.now(), TaskActionType.CREATE);
        assertTrue(isValid);
    }

    @Test
    public void isValidRateLimit_Failed_Throw_Exception() {
        String userId = "123123";
        Optional<Integer> rateLimitConfig = Optional.of(10);
        when(userConfigService.getRateLimitConfig(userId)).thenReturn(rateLimitConfig);
        Exception e = assertThrows(Exception.class, () ->
                rateLimitService.isValidRateLimit(userId, LocalDate.now(), TaskActionType.CREATE));
        assertNotNull(e);
    }

    @Test
    public void increaseCounterAnCheckRateLimit_Success_With_Valid_Counter() {
        String userId = "123123";
        Optional<Integer> rateLimitConfig = Optional.of(10);
        LocalDate now = LocalDate.now();
        when(redisTemplate.opsForHash()).thenReturn(hashOperations);
        when(hashOperations.increment(RedisUtils.getCounterAddedTaskPerUserPerDateKey(userId, now), userId, 1)).thenReturn(5L);
        when(userConfigService.getRateLimitConfig(userId)).thenReturn(rateLimitConfig);
        boolean isValid = rateLimitService.increaseCounterAnCheckRateLimit(userId, LocalDate.now());
        assertTrue(isValid);
    }

    @Test
    public void increaseCounterAnCheckRateLimit_Success_With_Invalid_Counter() {
        String userId = "123123";
        Optional<Integer> rateLimitConfig = Optional.of(3);
        LocalDate now = LocalDate.now();
        when(redisTemplate.opsForHash()).thenReturn(hashOperations);
        when(hashOperations.increment(RedisUtils.getCounterAddedTaskPerUserPerDateKey(userId, now), userId, 1)).thenReturn(5L);
        when(userConfigService.getRateLimitConfig(userId)).thenReturn(rateLimitConfig);
        boolean isValid = rateLimitService.increaseCounterAnCheckRateLimit(userId, LocalDate.now());
        assertFalse(isValid);
    }

    @Test
    public void decreaseCounterRateLimit_Success() {
        String userId = "123123";
        LocalDate now = LocalDate.now();

        when(redisTemplate.opsForHash()).thenReturn(hashOperations);
        when(hashOperations.increment(RedisUtils.getCounterAddedTaskPerUserPerDateKey(userId, now), userId, -1)).thenReturn(1L);
        long counter = rateLimitService.decreaseCounterRateLimit(userId, LocalDate.now());
        assertEquals(1, counter);
    }
}