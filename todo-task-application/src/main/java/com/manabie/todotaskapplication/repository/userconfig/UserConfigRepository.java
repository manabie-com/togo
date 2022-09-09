package com.manabie.todotaskapplication.repository.userconfig;

import com.manabie.todotaskapplication.data.model.UserConfig;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
@Repository
public interface UserConfigRepository extends JpaRepository<UserConfig, UUID> {
    UserConfig findByUserIdAndConfigType(String userId, String configType);
}
