package com.todo.ws.core.auth.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.todo.ws.core.auth.model.User;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {

}
