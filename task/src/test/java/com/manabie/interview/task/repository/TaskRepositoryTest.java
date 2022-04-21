package com.manabie.interview.task.repository;

import com.manabie.interview.task.model.Task;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;

import java.util.List;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

@DataJpaTest
class TaskRepositoryTest {

    @Autowired
    private TaskRepository test;
    @Test
    void checkFindDailyTaskByUserId() {
        test.save(new Task(1L, "hoavq", "19/04/2022", "unitest"));
        test.save(new Task(2L, "hoavq", "19/04/2022", "unitest"));
        test.save(new Task(3L, "hoavq", "18/04/2022", "unitest"));
        List<Task> tasks1904 = test.findDailyTaskByUserId("hoavq", "19/04/2022");
        List<Task> tasks1804 = test.findDailyTaskByUserId("hoavq", "18/04/2022");
        assertThat(tasks1904.size()).isEqualTo(2);
        assertThat(tasks1804.size()).isEqualTo(1);
    }
}