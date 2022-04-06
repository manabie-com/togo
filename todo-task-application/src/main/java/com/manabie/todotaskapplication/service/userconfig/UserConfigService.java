package com.manabie.todotaskapplication.service.userconfig;

import com.manabie.todotaskapplication.common.constant.ConfigType;
import com.manabie.todotaskapplication.data.model.UserConfig;

import java.util.Optional;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
public interface UserConfigService {
    
    Optional<UserConfig> getUserConfigByType(String userId, ConfigType configType);

    Optional<Integer> getRateLimitConfig(String userId);
}
