package com.manabie.togo.service;

import com.manabie.togo.dto.ToDoRequest;
import com.manabie.togo.exception.DailyLimitException;
import com.manabie.togo.exception.UserNotFoundException;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest
class ToDoServiceImplTest {

    @Autowired
    ToDoService toDoService;

    @Test
    public void addToDoTest(){
        assertEquals("To Do Task Added", toDoService.addToDo(ToDoRequest.builder()
                .userId("1111")
                .title("Title")
                .description("Description")
                .toDoDate("2022-01-01")
                .build()));
    }

    @Test
    public void addToDoUserNotFoundTest(){
        assertThrows(UserNotFoundException.class, () -> toDoService.addToDo(ToDoRequest.builder()
                .userId("6666")
                .title("Title")
                .description("Description")
                .toDoDate("2022-01-01")
                .build()));
    }

    @Test
    public void addToDoDailyLimitReachedTest(){
        assertThrows(DailyLimitException.class, () -> toDoService.addToDo(ToDoRequest.builder()
                .userId("5555")
                .title("Title")
                .description("Description")
                .toDoDate("2022-01-01")
                .build()));
    }

}