package com.todo.core.user.application.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.todo.core.commons.Messages;
import com.todo.core.commons.model.GenericResponse;
import com.todo.core.user.application.dto.CustomUserDetails;
import com.todo.core.user.application.dto.UserLoginDTO;
import com.todo.core.user.application.dto.UserRegistrationDTO;
import com.todo.core.user.exception.UserDoesNotExistException;
import com.todo.core.user.model.TodoUser;
import com.todo.core.user.repository.TodoUserRepository;
import com.todo.core.user.service.CustomUserDetailsService;
import com.todo.core.user.service.TodoUserService;
import com.todo.core.user.service.TokenProvider;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.test.context.support.WithMockUser;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.web.context.WebApplicationContext;

import java.nio.charset.Charset;
import java.util.Optional;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;
import static org.springframework.security.test.web.servlet.setup.SecurityMockMvcConfigurers.springSecurity;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@SpringBootTest
public class UserControllerTests {

    @Autowired
    private WebApplicationContext context;

    @Autowired
    private TokenProvider tokenProvider;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @MockBean
    private TodoUserService todoUserService;

    @MockBean
    private CustomUserDetailsService customUserDetailsService;

    @MockBean
    private TodoUserRepository todoUserRepository;


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

    @Test
    public void accessPublicLoginRouteOk() throws Exception {
        UserLoginDTO userLoginDTO = new UserLoginDTO(
            "userTest",
            "passTest"
        );

        String payload = new ObjectMapper().writeValueAsString(userLoginDTO);
        CustomUserDetails customUserDetails = mock(CustomUserDetails.class);
        when(customUserDetails.isAccountNonExpired()).thenReturn(true);
        when(customUserDetails.isAccountNonLocked()).thenReturn(true);
        when(customUserDetails.isEnabled()).thenReturn(true);
        when(customUserDetails.isCredentialsNonExpired()).thenReturn(true);
        when(customUserDetails.getUsername()).thenReturn(userLoginDTO.getUsername());
        when(customUserDetails.getPassword()).thenReturn(this.passwordEncoder.encode(userLoginDTO.getPassword()));

        when(customUserDetailsService.loadUserByUsername(userLoginDTO.getUsername()))
            .thenReturn(customUserDetails);


        this.mvc.perform(MockMvcRequestBuilders.post("/api/v1/users/login")
                .contentType(MediaType.APPLICATION_JSON)
                .characterEncoding(Charset.defaultCharset())
                .content(payload))
            .andExpect(status().isOk())
            .andDo(print());
    }

    @Test
    public void accessPublicLoginRouteBadCredentials() throws Exception {
        UserLoginDTO userLoginDTO = new UserLoginDTO(
            "userTest",
            "passTest"
        );

        String payload = new ObjectMapper().writeValueAsString(userLoginDTO);
        CustomUserDetails customUserDetails = mock(CustomUserDetails.class);
        when(todoUserRepository.findByUsername(anyString())).thenReturn(Optional.empty());

        when(customUserDetailsService.loadUserByUsername("userTest"))
            .thenThrow(UserDoesNotExistException.class);


        this.mvc.perform(MockMvcRequestBuilders.post("/api/v1/users/login")
                .contentType(MediaType.APPLICATION_JSON)
                .characterEncoding(Charset.defaultCharset())
                .content(payload))
            .andExpect(status().isOk())
            .andDo(print());
    }

    
}
