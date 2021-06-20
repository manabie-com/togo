package pro.datnt.manabie.task.controller;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.Captor;
import org.mockito.Mock;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.security.servlet.SecurityAutoConfiguration;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.http.MediaType;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.test.web.servlet.MockMvc;
import pro.datnt.manabie.task.controller.model.UserDTO;
import pro.datnt.manabie.task.model.UserDBO;
import pro.datnt.manabie.task.repository.UserRepository;
import pro.datnt.manabie.task.service.UserService;
import pro.datnt.manabie.task.service.security.JwtUtil;

import java.util.Optional;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.AdditionalMatchers.not;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.*;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(controllers = UserController.class, excludeAutoConfiguration = SecurityAutoConfiguration.class)
class UserControllerTest {

    @MockBean
    UserService userService;

    @MockBean
    JwtUtil jwtUtil;

    @Autowired
    MockMvc mockMvc;

    @Captor
    ArgumentCaptor<UserDTO> userDTOArgumentCaptor;

    @BeforeEach
    void setUp() {
    }

    @Test
    void loginWithRightPassword() throws Exception {
        when(userService.loginUser(any())).thenReturn(Optional.of("token"));
        mockMvc.perform(post("/user/login")
        .contentType(MediaType.APPLICATION_JSON)
        .content("{\"username\": \"tester\", \"password\": \"secret\"}"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.token").value("token"));

        verify(userService).loginUser(userDTOArgumentCaptor.capture());
        UserDTO value = userDTOArgumentCaptor.getValue();
        assertThat(value.getUsername()).isEqualTo("tester");
        assertThat(value.getPassword()).isEqualTo("secret");
    }
    @Test
    void loginWithWrongInfo() throws Exception {
        when(userService.loginUser(any())).thenReturn(Optional.empty());
        mockMvc.perform(post("/user/login")
        .contentType(MediaType.APPLICATION_JSON)
        .content("{\"username\": \"tester\", \"password\": \"secret\"}"))
                .andExpect(status().isBadRequest());

        verify(userService).loginUser(userDTOArgumentCaptor.capture());
        UserDTO value = userDTOArgumentCaptor.getValue();
        assertThat(value.getUsername()).isEqualTo("tester");
        assertThat(value.getPassword()).isEqualTo("secret");
    }

    @Test
    void register() throws Exception {
        mockMvc.perform(post("/user/register")
                .contentType(MediaType.APPLICATION_JSON)
                .content("{\"username\": \"tester1\", \"password\": \"secret\"}"))
                .andExpect(status().isAccepted());
        verify(userService).createUser(userDTOArgumentCaptor.capture());
        UserDTO userDTO = userDTOArgumentCaptor.getValue();
        assertThat(userDTO.getUsername()).isEqualTo("tester1");
        assertThat(userDTO.getPassword()).isEqualTo("secret");
    }

    @Test
    void registerWrongData() throws Exception {
        doThrow(new DataIntegrityViolationException("test")).when(userService).createUser(any());
        mockMvc.perform(post("/user/register")
                .contentType(MediaType.APPLICATION_JSON)
                .content("{\"username\": \"tester1\", \"password\": \"secret\"}"))
                .andExpect(status().isBadRequest());
        verify(userService).createUser(userDTOArgumentCaptor.capture());
        UserDTO userDTO = userDTOArgumentCaptor.getValue();
        assertThat(userDTO.getUsername()).isEqualTo("tester1");
        assertThat(userDTO.getPassword()).isEqualTo("secret");
    }
}