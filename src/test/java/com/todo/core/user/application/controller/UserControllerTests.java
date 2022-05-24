package com.todo.core.user.application.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.todo.core.commons.Messages;
import com.todo.core.commons.model.GenericResponse;
import com.todo.core.user.application.dto.UserRegistrationDTO;
import com.todo.core.user.service.TodoUserService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultHandlers;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.web.context.WebApplicationContext;

import java.nio.charset.Charset;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.springframework.security.test.web.servlet.setup.SecurityMockMvcConfigurers.springSecurity;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@SpringBootTest
public class UserControllerTests {

    @Autowired
    private WebApplicationContext context;

    @MockBean
    private TodoUserService todoUserService;


    private MockMvc mvc;

    @BeforeEach
    public void setup() {
        this.mvc = MockMvcBuilders
            .webAppContextSetup(context)
            .apply(springSecurity())
            .build();
    }


    @Test
    public void accessPublicRouteOk() throws Exception {
        this.mvc.perform(MockMvcRequestBuilders.get("/api/v1/version").contentType(MediaType.APPLICATION_JSON))
            .andExpect(status().isOk())
            .andDo(print());
    }

    @Test
    public void accessPublicRegisterRouteOk() throws Exception {

        UserRegistrationDTO userRegistrationDTO = new UserRegistrationDTO(
            "userTest",
            "passTest",
            12
        );

        String payload = new ObjectMapper().writeValueAsString(userRegistrationDTO);
        when(this.todoUserService.createUser(any(UserRegistrationDTO.class)))
            .thenReturn(new GenericResponse(true, Messages.USER_CREATE_SUCCESSFUL.getContent()));

        this.mvc.perform(MockMvcRequestBuilders.post("/api/v1/users/create")
            .contentType(MediaType.APPLICATION_JSON)
                .characterEncoding(Charset.defaultCharset())
                .content(payload))
            .andExpect(status().isOk())
            .andDo(print());
    }

    
}
