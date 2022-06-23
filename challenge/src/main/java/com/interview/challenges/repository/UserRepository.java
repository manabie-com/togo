package com.interview.challenges.repository;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import com.interview.challenges.domain.User;

@Repository
public interface UserRepository extends CrudRepository<User, Long>{
	User findByUsername(String username);
}
