package com.manabie.assignment.repositories;

import org.springframework.data.repository.CrudRepository;

import com.manabie.assignment.models.User;

public interface UserRepository extends CrudRepository<User, Integer> {

}
