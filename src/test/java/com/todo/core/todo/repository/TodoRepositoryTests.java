package com.todo.core.todo.repository;

import com.todo.core.todo.model.Todo;
import com.todo.core.user.model.TodoUser;
import com.todo.core.user.repository.TodoUserRepository;
import org.assertj.core.util.Lists;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import javax.persistence.EntityManager;
import javax.sql.DataSource;
import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;

@ExtendWith(SpringExtension.class)
@DataJpaTest
public class TodoRepositoryTests {

    @Autowired
    private DataSource dataSource;
    @Autowired private JdbcTemplate jdbcTemplate;
    @Autowired private EntityManager entityManager;
    @Autowired private TodoRepository todoRepository;
    @Autowired private TodoUserRepository todoUserRepository;

    @Test
    void doCheckInjectectedComponents() {
        assertThat(todoRepository).isNotNull();
    }

    @Test
    void doSaveAndListAllFromUser() {
        todoUserRepository
            .save(new
                TodoUser("usertest", "passtest", 5L)
            );

        final Long id = todoUserRepository.findByUsername("usertest")
            .get()
            .getId();

        todoRepository.saveAll(Lists.list(
                new Todo(
                    "not-completed",
                    "task-test",
                    id
                ),
                new Todo(
                    "not-completed",
                    "task-test-2",
                    id
                ),
                new Todo(
                    "not-completed",
                    "task-test-3",
                    77L
                )

        ));

        List<Todo> todoList = todoRepository.findAllByTodoUserId(id);
        assertThat(todoList.size()).isEqualTo(2);
    }
}
