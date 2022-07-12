package com.todo.ws.todo.repository;

import com.todo.ws.todo.model.Todo;
import com.todo.ws.user.model.TodoUser;
import com.todo.ws.user.repository.TodoUserRepository;
import org.assertj.core.util.Lists;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import javax.persistence.EntityManager;
import javax.sql.DataSource;
import java.time.LocalDate;
import java.time.ZoneId;
import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

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

    @BeforeEach
    void doBeforeEach() {
        todoRepository.deleteAll();
        todoUserRepository.deleteAll();
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

        final Pageable pageable = mock(Pageable.class);
        when(pageable.getOffset()).thenReturn(0L);
        when(pageable.getPageNumber()).thenReturn(1);
        when(pageable.getPageSize()).thenReturn(3);
        when(pageable.getSort()).thenReturn(Sort.by("dateCreated"));


        todoRepository.saveAll(Lists.list(
                new Todo(
                    "not-completed",
                    "task-test",
                    id,
                    LocalDate.now(ZoneId.of("Asia/Manila"))
                ),
                new Todo(
                    "not-completed",
                    "task-test-2",
                    id,
                    LocalDate.now(ZoneId.of("Asia/Manila"))
                ),
                new Todo(
                    "not-completed",
                    "task-test",
                    id,
                    LocalDate.now(ZoneId.of("Asia/Manila"))
                ),
                new Todo(
                    "not-completed",
                    "task-test-2",
                    id,
                    LocalDate.now(ZoneId.of("Asia/Manila"))
                ),
                new Todo(
                    "not-completed",
                    "task-test",
                    id,
                    LocalDate.now(ZoneId.of("Asia/Manila"))
                ),
                new Todo(
                    "not-completed",
                    "task-test-2",
                    id,
                    LocalDate.now(ZoneId.of("Asia/Manila"))
                ),
                new Todo(
                    "not-completed",
                    "task-test-3",
                    77L,
                    LocalDate.now(ZoneId.of("Asia/Manila"))
                )

        ));

        Page<Todo> todoList = todoRepository.findAllByTodoUserId(id, pageable);
        assertThat(todoList.getNumberOfElements()).isEqualTo(3);
        assertThat(todoList.getTotalElements()).isEqualTo(6);
        assertThat(todoList.getTotalPages()).isEqualTo(2);
    }

    @Test
    public void retrieveByDateCreatedAndUserId() {
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
                id,
                LocalDate.now(ZoneId.of("Asia/Manila"))
            ),
            new Todo(
                "not-completed",
                "task-test-2x",
                id,
                LocalDate.of(2002, 12,12)
            ),
            new Todo(
                "not-completed",
                "task-test-2",
                id,
                LocalDate.of(2002, 12,12)
            ),
            new Todo(
                "not-completed",
                "task-test-3",
                77L,
                LocalDate.now(ZoneId.of("Asia/Manila"))
            )

        ));

        final int countOfToday =
            todoRepository.countByDateCreatedAndTodoUserId(LocalDate.now(ZoneId.of("Asia/Manila")), id);
        final int countOfOtherDay =
            todoRepository.countByDateCreatedAndTodoUserId(LocalDate.of(2002, 12, 12), id);
        assertThat(countOfToday).isEqualTo(1);
        assertThat(countOfOtherDay).isEqualTo(2);
    }
}
