package com.manabie.interview.task.repository;

import com.manabie.interview.task.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface UserRepository extends JpaRepository<User, String> {
    @Query("SELECT u FROM User u WHERE u.uid = ?1")
    Optional<User> findUserById(String uid);
    @Query("SELECT u FROM User u WHERE u.uid = ?1 AND u.upassword = ?2")
    Optional<User> findUserByIdAndPassword(String uid, String password);
}
