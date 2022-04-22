package com.manabie.interview.task.service;

import com.manabie.interview.task.model.Task;
import com.manabie.interview.task.model.User;
import com.manabie.interview.task.model.UserRole;
import com.manabie.interview.task.repository.TaskRepository;
import com.manabie.interview.task.repository.UserRepository;
import com.manabie.interview.task.response.APIResponse;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.http.HttpStatus;

import java.util.List;
import java.util.Optional;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.verify;
import static org.mockito.BDDMockito.given;
import static org.mockito.Mockito.verify;
@ExtendWith(MockitoExtension.class)
class TaskServiceTest {

    private TaskService test;

    @Mock
    private TaskRepository taskRepository;

    @Mock
    private UserRepository userRepository;

    @BeforeEach
    void setup(){
        test = new TaskService(taskRepository, userRepository);
    }

    @AfterEach
    void teardown(){

    }

    @Test
    void checkGetAllTasks() {
        test.getTasks();
        verify(taskRepository).findAll();
    }

    @Test
    void checkRegisterNewTaskUserNotExisted() {
        Task task = new Task(1L, "nhanptt", "19/04/2022", "unitest");
//        given(userRepository.findUserById("nhanptt")).willReturn(true);
        APIResponse res = new APIResponse(String.format("User %s doesn't exist. Can't not assign task", task.getUserUid()), HttpStatus.OK);
        assertThat(test.registerNewTask(task)).hasToString(res.toString());
    }

    @Test
    void checkRegisterNewTaskWithLimitedTask() {
        Task task = new Task(1L, "hoavq", "19/04/2022", "unitest");
        User hoavq = new User("hoavq", "1234",1, UserRole.USER);
        List<Task> list =  List.of(task);
        given(userRepository.findUserById(hoavq.getUid())).willReturn(Optional.of(hoavq));
        given(taskRepository.findDailyTaskByUserId(task.getUserUid(), task.getCreatedDate())).willReturn(list);
        APIResponse res = new APIResponse(String.format("User %s reaches limit daily task. Can't not assign task", task.getUserUid()), HttpStatus.OK);
        assertThat(test.registerNewTask(task)).hasToString(res.toString());
    }

    @Test
    void checkRegisterNewTaskOke() {
        Task task = new Task(1L, "hoavq", "19/04/2022", "unitest");
        User hoavq = new User("hoavq", "1234",2, UserRole.USER);
        List<Task> list =  List.of(task);
        given(userRepository.findUserById(hoavq.getUid())).willReturn(Optional.of(hoavq));
        given(taskRepository.findDailyTaskByUserId(task.getUserUid(), task.getCreatedDate())).willReturn(list);
        APIResponse res = new APIResponse("Assigned Successfully", HttpStatus.OK);
        assertThat(test.registerNewTask(task)).hasToString(res.toString());
    }
}