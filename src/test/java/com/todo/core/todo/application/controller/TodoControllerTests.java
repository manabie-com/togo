package com.todo.core.todo.application.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.todo.core.todo.application.dto.TodoDTO;
import com.todo.core.todo.model.Todo;
import com.todo.core.todo.service.TodoServiceImpl;
import com.todo.core.user.model.TodoUser;
import com.todo.core.user.repository.TodoUserRepository;
import com.todo.core.user.service.TokenProvider;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.http.MediaType;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.config.annotation.authentication.configuration.AuthenticationConfiguration;
import org.springframework.security.core.Authentication;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.web.context.WebApplicationContext;

import java.util.Optional;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertWith;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;
import static org.springframework.security.test.web.servlet.setup.SecurityMockMvcConfigurers.springSecurity;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@SpringBootTest
public class TodoControllerTests {

    @Autowired
    private WebApplicationContext context;

    @Autowired
    private TodoServiceImpl todoService;

    @Autowired
    private TodoUserRepository todoUserRepository;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @Autowired
    private AuthenticationConfiguration authenticationConfiguration;

    @Autowired
    private TokenProvider tokenProvider;

    private MockMvc mvc;

    @BeforeEach
    public void setup() {
        this.mvc = MockMvcBuilders
            .webAppContextSetup(context)
            .apply(springSecurity())
            .build();
        this.todoUserRepository.deleteAll();
    }

    @Test
    public void testGetAllTodoOfUserOk() throws Exception {
        doSaveMockUser();

        final Pageable pageable = mock(Pageable.class);

        when(pageable.getOffset()).thenReturn(0L);
        when(pageable.getPageNumber()).thenReturn(1);
        when(pageable.getPageSize()).thenReturn(3);
        when(pageable.getSort()).thenReturn(Sort.by("dateCreated"));

        Authentication authentication = authenticationConfiguration.getAuthenticationManager().authenticate(
            new UsernamePasswordAuthenticationToken(
                "usertest",
                "passtest"
            )
        );
        String jwt = tokenProvider.generateToken(authentication);

        this.mvc.perform(MockMvcRequestBuilders.get("/api/v1/todos/list")
                .header("Authorization", "Bearer " + jwt)
                .queryParam("page", "0")
                .queryParam("size", "3")
                .contentType(MediaType.APPLICATION_JSON))
            .andExpect(status().isOk())
            .andDo(print());
    }

    @Test
    public void createTodoForUserOk() throws Exception {
        doSaveMockUser();
        Authentication authentication = authenticationConfiguration.getAuthenticationManager().authenticate(
            new UsernamePasswordAuthenticationToken(
                "usertest",
                "passtest"
            )
        );
        String jwt = tokenProvider.generateToken(authentication);


        TodoDTO todoDTO = new TodoDTO();
        todoDTO.setTodoUserId(1L);
        todoDTO.setStatus("not-completed");
        todoDTO.setTask("Task-test 1");

        String payload = new ObjectMapper().writeValueAsString(todoDTO);

        this.mvc.perform(MockMvcRequestBuilders.get("/api/v1/todos/create")
                .header("Authorization", "Bearer " + jwt)
                .queryParam("page", "0")
                .queryParam("size", "3")
                .contentType(MediaType.APPLICATION_JSON).content(payload))
            .andExpect(status().isOk())
            .andDo(print());
    }

    private void doSaveMockUser() {

        todoUserRepository
            .save(new
                TodoUser("usertest", this.passwordEncoder.encode("passtest"), 5L)
            );

        final Optional<TodoUser> user = todoUserRepository.findByUsername("usertest");

        assertThat(user.isPresent()).isTrue();
        assertWith(user.get().getUsername()).isEqualTo("usertest");
    }

}
