package com.api.todo.services;

import static org.assertj.core.api.Assertions.assertThat;

import java.time.LocalDate;
import java.time.format.DateTimeFormatter;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import com.api.todo.entities.Task;
import com.api.todo.entities.User;
import com.api.todo.repositories.TaskRepository;
import com.api.todo.repositories.UserRepository;

@SpringBootTest
public class TodoServiceTest {
    @Autowired
    private TaskRepository taskRepository;
    @Autowired
    private UserRepository userRepository;
    @Autowired
    private TodoService todoService;
    private User userSaved;

    @BeforeEach
    void setInitial() {
        User user = new User("User A Test", 4);
        userSaved = userRepository.save(user);
    }

    @AfterEach
    void cleanInitial() {
        taskRepository.deleteAll();
        System.out.println("Deleted all data of task table");
        userRepository.deleteAll();
        System.out.println("Deleted all data of user table");
    }

    @DisplayName("Create task success")
    @Test
    void createTask_success() {
        Task task = new Task("title", "description", userSaved.getId());

        Task task_result = todoService.createTask(task);

        assertThat(task_result.getTitle()).isEqualTo(task.getTitle());
        assertThat(task_result.getDescription()).isEqualTo(task.getDescription());
        assertThat(task_result.getUserId()).isEqualTo(task.getUserId());
    }

    @DisplayName("Count tasks of user")
    @Test
    void countTaskOfOneUser_return_right() {
        LocalDate dateObj = LocalDate.now();
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd");
        String currentDate = dateObj.format(formatter);

        Task task = new Task("title", "description", userSaved.getId());
        Task task1 = new Task("title1", "description", userSaved.getId());
        Task task2 = new Task("title2", "description", userSaved.getId());

        todoService.createTask(task);
        todoService.createTask(task1);
        todoService.createTask(task2);
        int count = todoService.countTaskOfOneUser(userSaved.getId(), currentDate);

        assertThat(count).isEqualTo(3);
    }
}
