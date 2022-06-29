package com.manabie.assignment;

import com.manabie.assignment.controllers.TaskController;
import com.manabie.assignment.repositories.TaskRepository;
import com.manabie.assignment.repositories.TasksPerDayRepository;
import com.manabie.assignment.repositories.UserRepository;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.context.ApplicationContext;

@SpringBootTest
class AssignmentApplicationTests {

    @Autowired
    private ApplicationContext context;

    @Test
    public void testBeans() {
        Assertions.assertNotNull(context.getBean(TaskController.class));
        Assertions.assertNotNull(context.getBean(TaskRepository.class));
        Assertions.assertNotNull(context.getBean(TasksPerDayRepository.class));
        Assertions.assertNotNull(context.getBean(UserRepository.class));
    }
}
