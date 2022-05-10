package com.todo;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.todo.repository.TodoTaskRepository;
import net.bytebuddy.utility.RandomString;
import org.junit.jupiter.api.Test;
import org.junit.runner.RunWith;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.security.test.context.support.WithMockUser;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@RunWith(SpringRunner.class)
@SpringBootTest
@AutoConfigureMockMvc
public class TodoTaskControllerTest {
    @Autowired
    MockMvc mockMvc;
    @Autowired
    ObjectMapper mapper;

    @MockBean
    TodoTaskRepository todoTaskRepository;

    String taskContent1 = RandomString.make(200);
    String taskContent2 = RandomString.make(200);
    String taskContent3 = RandomString.make(200);
    String taskContent4 = RandomString.make(200);

    @Test
    public void findAll_success() throws Exception {
//        Mockito.when(todoTaskRepository.findTodoTaskByAppAccount_Username("aa@aaaa")).thenReturn(records);

//        mockMvc.perform(MockMvcRequestBuilders
//                        .get("/tasks")
//                        .contentType(MediaType.APPLICATION_JSON))
//                .andExpect(status().isOk());
//                .andExpect(jsonPath("$", hasSize(3)))
//                .andExpect(jsonPath("$[2].name", is("Jane Doe")));

        mockMvc.perform(MockMvcRequestBuilders.post("/tasks").header("authen", "asdasf")).andExpect(status().isForbidden());

    }
}
