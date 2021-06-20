package pro.datnt.manabie.task.service;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.*;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import pro.datnt.manabie.task.controller.model.UserDTO;
import pro.datnt.manabie.task.model.UserDBO;
import pro.datnt.manabie.task.repository.UserRepository;
import pro.datnt.manabie.task.service.security.JwtUtil;

import java.util.Optional;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.AdditionalMatchers.not;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class UserServiceTest {
    @Mock
    UserRepository userRepository;

    @InjectMocks
    UserService userService;

    @Spy
    PasswordEncoder passwordEncoder = new BCryptPasswordEncoder();

    @Captor
    ArgumentCaptor<UserDBO> userDBOArgumentCaptor;

    @Mock
    JwtUtil jwtUtil;

    UserDBO user = new UserDBO();

    @BeforeEach
    public void setup() {
        user.setUsername("tester");
        user.setPassword(passwordEncoder.encode("secret"));
    }

    @Test
    void getUserId() {
        // Setup
        UserDBO userDBO = new UserDBO();
        userDBO.setId(1L);
        when(userRepository.findByUsername("tester")).thenReturn(Optional.of(userDBO));
        // Execution
        Optional<Long> userId = userService.getUserId("tester");
        // Verification
        assertThat(userId.get()).isEqualTo(1L);
    }

    @Test
    void loginUserWithEmpty() {
        when(userRepository.findByUsername(any())).thenReturn(Optional.empty());
        // Execution
        Optional<String> token = userService.loginUser(new UserDTO());
        // Verification
        assertThat(token.isEmpty()).isTrue();
    }

    @Test
    void loginUserWithWrongPassword() {
        when(userRepository.findByUsername("tester")).thenReturn(Optional.of(user));
        // Execution
        UserDTO user = new UserDTO();
        user.setUsername("tester");
        user.setPassword("1234");
        Optional<String> token = userService.loginUser(user);
        // Verification
        assertThat(token.isEmpty()).isTrue();
    }

    @Test
    void loginUserWithWrongUser() {
        when(userRepository.findByUsername(not(eq("tester")))).thenReturn(Optional.empty());
        // Execution
        UserDTO user = new UserDTO();
        user.setUsername("tester1234");
        user.setPassword("1234");
        Optional<String> token = userService.loginUser(user);
        // Verification
        assertThat(token.isEmpty()).isTrue();
    }
    @Test
    void loginUserWithRightInfo() {
        when(userRepository.findByUsername("tester")).thenReturn(Optional.of(user));
        when(jwtUtil.generateToken("tester")).thenReturn("token");
        UserDTO userDTO = new UserDTO();
        userDTO.setUsername("tester");
        userDTO.setPassword("secret");
        // Execution
        Optional<String> token = userService.loginUser(userDTO);
        // Verification
        assertThat(token.get()).isEqualTo("token");
    }

    @Test
    void createUser() {
        UserDTO userDTO = new UserDTO();
        userDTO.setUsername("tester");
        userDTO.setPassword("secret");
        // Execution
        userService.createUser(userDTO);
        // Verification
        verify(userRepository).save(userDBOArgumentCaptor.capture());
        UserDBO value = userDBOArgumentCaptor.getValue();
        assertThat(value.getUsername()).isEqualTo("tester");
        assertThat(passwordEncoder.matches("secret", value.getPassword())).isTrue();
    }
}