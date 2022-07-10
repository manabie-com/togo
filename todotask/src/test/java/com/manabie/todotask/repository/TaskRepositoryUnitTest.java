package com.manabie.todotask.repository;

import com.manabie.todotask.entity.Task;
import com.manabie.todotask.entity.UserDailyLimit;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.jdbc.AutoConfigureTestDatabase;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.annotation.DirtiesContext;

import java.time.ZonedDateTime;
import java.util.Optional;

@SpringBootTest
@AutoConfigureTestDatabase
@DirtiesContext(classMode = DirtiesContext.ClassMode.BEFORE_EACH_TEST_METHOD)
class TaskRepositoryUnitTest {

    @Autowired
    TaskRepository taskRepository;
    @Test
    public void addAndFind(){
        int userId = 1;
        ZonedDateTime targetDate = ZonedDateTime.parse("2022-07-08T00:00:00Z");

        Task task = new Task();
        task.setName("test");
        task.setUserId(userId);
        task.setDescription("test");
        task.setTargetDate(targetDate);
        Task t = taskRepository.save(task);
        System.out.println(t);
        assert(t.getId().equals(1));

        Optional<Task> indbOpt = taskRepository.findById(t.getId());
        assert (indbOpt.isPresent());
        Task indb = indbOpt.get();

        assert (indbOpt.get().getUserId().equals(userId));
        assert (indbOpt.get().getName().equals("test"));
        assert (indbOpt.get().getDescription().equals("test"));
        assert (indbOpt.get().getTargetDate().toInstant().equals(targetDate.toInstant()));

    }

    @Test
    void countTaskByUserIdAndTargetDate() {
        int userId = 1;
        int userId2 = 2;
        ZonedDateTime targetDate = ZonedDateTime.parse("2022-07-08T00:00:00Z");
        ZonedDateTime targetDate2 = ZonedDateTime.parse("2022-07-09T00:00:00Z");

        Task task = new Task();
        task.setName("test");
        task.setUserId(userId);
        task.setDescription("test");
        task.setTargetDate(targetDate);
        Task t = taskRepository.save(task);
        System.out.println(t);
        assert(t.getId().equals(1));

        Task task2 = new Task();
        task2.setName("test");
        task2.setUserId(userId);
        task2.setDescription("test");
        task2.setTargetDate(targetDate);
        Task t2 = taskRepository.save(task2);
        System.out.println(t2);
        assert(t2.getId().equals(2));

        Task task3 = new Task();
        task3.setName("test");
        task3.setUserId(userId2);
        task3.setDescription("test");
        task3.setTargetDate(targetDate);
        Task t3 = taskRepository.save(task3);
        System.out.println(t3);
        assert(t3.getId().equals(3));

        Task task4 = new Task();
        task4.setName("test");
        task4.setUserId(userId);
        task4.setDescription("test");
        task4.setTargetDate(targetDate2);
        Task t4 = taskRepository.save(task4);
        System.out.println(t4);
        assert(t4.getId().equals(4));

        int count = taskRepository.countTaskByUserIdAndTargetDate(userId, targetDate, targetDate.plusDays(1));
        assert (count == 2);
    }
}