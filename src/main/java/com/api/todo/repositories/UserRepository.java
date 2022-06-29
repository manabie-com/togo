package com.api.todo.repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.api.todo.entities.User;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {}

