package com.manabie.interview.task.repository;

import com.manabie.interview.task.model.User;
import com.manabie.interview.task.model.UserRole;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;

import java.util.Optional;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;
import static org.junit.jupiter.api.Assertions.*;

@DataJpaTest
class UserRepositoryTest {

    @Autowired
    private UserRepository test;

    @AfterEach
    void tearDown(){
        test.deleteAll();
    }

    @Test
    void checkFindUserById() {
        User hoavq = new User("hoavq", "1234",5, UserRole.USER);
        test.save(hoavq);
        Optional<User> user = test.findUserById("hoavq");
        assertThat(user.get()).isEqualTo(hoavq);
    }

    @Test
    void checkFindUserByIdAndPassword() {
        User hoavq = new User("hoavq", "1234",5, UserRole.USER);
        test.save(hoavq);
        Optional<User> user1 = test.findUserByIdAndPassword("hoavq", "1234");
        assertThat(user1.get()).isEqualTo(hoavq);

        Optional<User> user2 = test.findUserByIdAndPassword("hoavq", "12345");
        assertThat(user2).isEmpty();
    }
}