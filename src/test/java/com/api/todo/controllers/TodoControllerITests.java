package com.api.todo.controllers;

import static org.hamcrest.CoreMatchers.is;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.ResultActions;

import com.api.todo.entities.User;
import com.api.todo.repositories.TaskRepository;
import com.api.todo.repositories.UserRepository;
import com.api.todo.request.RequestTaskEntity;
import com.fasterxml.jackson.databind.ObjectMapper;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@AutoConfigureMockMvc
public class TodoControllerITests {

    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private UserRepository userRepository;
    @Autowired
    private TaskRepository taskRepository;

    @Autowired
    private ObjectMapper objectMapper;
    
    private User userSaved;

    @BeforeEach
    void setup(){
    	// give
    	User user = new User("User A Test", 4);
        userSaved = userRepository.save(user);
    }
    @AfterEach
    void tearDown(){
    	// clean
    	taskRepository.deleteAll();
    	userRepository.deleteAll();
    }

    @Test
    public void givenEmployeeObject_whenCreateTask_thenReturnSavedTask() throws Exception{

        RequestTaskEntity requestTask = new RequestTaskEntity("title_test", "description_test", userSaved.getId());

        // when - action or behavior that we are going test
        ResultActions response = mockMvc.perform(post("/api/todo/tasks")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(requestTask)));

        // then - verify the result or output using assert statements
        response.andDo(print()).
                andExpect(status().isCreated())
                .andExpect(jsonPath("$.title",
                        is(requestTask.getTitle())))
                .andExpect(jsonPath("$.description",
                        is(requestTask.getDescription())))
                .andExpect(jsonPath("$.userId",
                        is(Integer.parseInt(""+requestTask.getUserId()))));

    }
    
    @Test
    public void givenEmployeeObject_whenCreateTask_thenReturn404() throws Exception{

        RequestTaskEntity requestTask = new RequestTaskEntity("title_test", "description_test", userSaved.getId());

        // when - action or behaviour that we are going test
        ResultActions response = mockMvc.perform(post("/api/todo/task")
                .contentType(MediaType.APPLICATION_JSON)
                .content(objectMapper.writeValueAsString(requestTask)));

        // then - verify the result or output using assert statements
        response.andExpect(status().isNotFound()).andDo(print());

    }
}
