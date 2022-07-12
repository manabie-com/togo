package com.todo.ws.user.repository;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import com.todo.ws.user.model.TodoUser;

import javax.persistence.EntityManager;
import javax.sql.DataSource;
import java.util.Optional;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertWith;

@ExtendWith(SpringExtension.class)
@DataJpaTest
public class TodoUserRepositoryTests {

    @Autowired private DataSource dataSource;
    @Autowired private JdbcTemplate jdbcTemplate;
    @Autowired private EntityManager entityManager;
    @Autowired private TodoUserRepository todoUserRepository;

    @Test
    void doCheckInjectectedComponents() {
        assertThat(todoUserRepository).isNotNull();
    }

    @Test
    void doTrySaveAndRetrieve() {
        todoUserRepository
            .save(new
                TodoUser("usertest", "passtest", 5L)
            );

        final Optional<TodoUser> user = todoUserRepository.findByUsername("usertest");

        assertThat(user.isPresent()).isTrue();
        assertWith(user.get().getUsername()).isEqualTo("usertest");
    }

}
