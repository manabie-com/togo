package com.manabie.togo.repository;

import com.manabie.togo.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

/**
 * Data access layer for user
 * @author mupmup
 */
@Repository
public interface UserRepository extends JpaRepository<User, String> {

    User findByUsername(String username);
}
