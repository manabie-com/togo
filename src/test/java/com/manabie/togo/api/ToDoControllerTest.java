package com.manabie.togo.api;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.manabie.togo.dto.ToDoRequest;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;


@SpringBootTest
@AutoConfigureMockMvc
class ToDoControllerTest {

    @Autowired
    private MockMvc mvc;

    @Autowired
    private ObjectMapper objectMapper;

    @Test
    public void addToDoTest() throws Exception {
        mvc.perform(post("/todo/add").contentType(MediaType.APPLICATION_JSON_VALUE)
                .content(objectMapper.writeValueAsString(ToDoRequest.builder()
                        .userId("1111")
                        .title("Title")
                        .description("Description")
                        .toDoDate("2022-01-01")
                        .build())))
                .andExpect(status().isOk());
    }

    @Test
    public void addToDoUserNotFoundTest() throws Exception {
        mvc.perform(post("/todo/add").contentType(MediaType.APPLICATION_JSON_VALUE)
                .content(objectMapper.writeValueAsString(ToDoRequest.builder()
                        .userId("6666")
                        .title("Title")
                        .description("Description")
                        .toDoDate("2022-01-01")
                        .build())))
                .andExpect(status().isNotFound());
    }

    @Test
    public void addToDoDailyLimitReachedTest() throws Exception {
        mvc.perform(post("/todo/add").contentType(MediaType.APPLICATION_JSON_VALUE)
                .content(objectMapper.writeValueAsString(ToDoRequest.builder()
                        .userId("5555")
                        .title("Title")
                        .description("Description")
                        .toDoDate("2022-01-01")
                        .build())))
                .andExpect(status().isTooManyRequests());
    }

}