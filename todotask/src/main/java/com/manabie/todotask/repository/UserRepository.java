package com.manabie.todotask.repository;

import com.manabie.todotask.entity.UserDailyLimit;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import javax.persistence.Entity;

@Repository
public interface UserRepository extends CrudRepository<UserDailyLimit, Integer> {
}
