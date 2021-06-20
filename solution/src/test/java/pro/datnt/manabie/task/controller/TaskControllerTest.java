package pro.datnt.manabie.task.controller;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import org.checkerframework.checker.nullness.Opt;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.security.servlet.SecurityAutoConfiguration;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import pro.datnt.manabie.task.model.TaskDBO;
import pro.datnt.manabie.task.service.TaskService;
import pro.datnt.manabie.task.service.UserService;
import pro.datnt.manabie.task.service.security.JwtUtil;

import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(value = TaskController.class, excludeAutoConfiguration = SecurityAutoConfiguration.class)
class TaskControllerTest {
    @MockBean
    JwtUtil jwtUtil;

    @Autowired
    MockMvc mockMvc;

    @MockBean
    TaskService taskService;

    @MockBean
    UserService userService;

    @BeforeEach
    void setUp() {
        Claims tester = Jwts.claims().setSubject("tester");
        when(jwtUtil.parseToken("token")).thenReturn(tester);
        when(userService.getUserId("tester")).thenReturn(Optional.of(99L));
    }

    @Test
    void createTaskWithInvalidToken() throws Exception {
        mockMvc.perform(post("/task/add")
            .contentType(MediaType.APPLICATION_JSON)
                .header("Authorization", "invalidToken")
            .content("{}"))
                .andExpect(status().isForbidden());
    }

    @Test
    void createTaskWithValidToken() throws Exception {
        when(taskService.createTask("task 1", 99L)).thenReturn(new TaskDBO());
        mockMvc.perform(post("/task/add")
            .contentType(MediaType.APPLICATION_JSON)
                .header("Authorization", "token")
            .content("{\"content\": \"task 1\"}"))
                .andExpect(status().isAccepted());
        verify(taskService).createTask("task 1", 99L);
    }
}