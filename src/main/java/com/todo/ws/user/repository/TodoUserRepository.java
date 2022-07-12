package com.todo.ws.user.repository;

import com.todo.ws.user.model.TodoUser;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.Optional;

public interface TodoUserRepository extends JpaRepository<TodoUser, Long> {

    @Query("SELECT f FROM TodoUser f WHERE f.username = :username")
    Optional<TodoUser> findByUsername(@Param("username") String username);
}
