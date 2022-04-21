package com.manabie.interview.task.service;


import com.manabie.interview.task.model.User;
import com.manabie.interview.task.model.UserRole;
import com.manabie.interview.task.repository.UserRepository;
import com.manabie.interview.task.response.APIResponse;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.ArgumentCaptor;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;

import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.http.HttpStatus;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.Optional;

import static org.assertj.core.api.AssertionsForClassTypes.*;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.BDDMockito.given;
import static org.mockito.Mockito.verify;

@ExtendWith(MockitoExtension.class)
class UserServiceTest {


    private UserService test;
    @Mock
    private UserRepository repo;
    @Mock
    BCryptPasswordEncoder bCryptPasswordEncoder;




    @BeforeEach
    void setUp(){
        test = new UserService(repo, bCryptPasswordEncoder);
    }

    @AfterEach
    void tearDown() throws Exception {
    }

    @Test
    void checkRegisterNewUserNonExisted() {
        User hoavq = new User("hoavq", "1234",5, UserRole.USER);
        test.registerNewUser(hoavq);
        ArgumentCaptor<User> userArgumentCaptor = ArgumentCaptor.forClass(User.class);
        verify(repo).save(userArgumentCaptor.capture());
        User captureUser = userArgumentCaptor.getValue();
        assertThat(captureUser).isEqualTo(hoavq);
    }

    @Test
    void checkRegisterNewUserExisted() {
        User hoavq = new User("hoavq", "1234",5, UserRole.USER);
        given(repo.existsById(hoavq.getUid())).willReturn(true);
        APIResponse res = new APIResponse("UserID existed", HttpStatus.OK);
        assertThat(test.registerNewUser(hoavq)).hasToString(res.toString());
    }

    @Test
    void checkDeleteUserExisted() {
        User hoavq = new User("hoavq", "1234",5, UserRole.USER);
        given(repo.existsById(hoavq.getUid())).willReturn(true);
        APIResponse res = new APIResponse("Successfully Remove", HttpStatus.OK);
        assertThat(test.deleteUser(hoavq.getUid())).hasToString(res.toString());
    }

    @Test
    void checkDeleteUserNonExisted() {
        User hoavq = new User("hoavq", "1234",5, UserRole.USER);
        given(repo.existsById(hoavq.getUid())).willReturn(false);
        APIResponse res = new APIResponse("UserID not existed", HttpStatus.OK);
        assertThat(test.deleteUser(hoavq.getUid())).hasToString(res.toString());
    }

    @Test
    void checkLoadUserByUsernameExisted() {
        User xxx = new User("hoavq", "1234",5, UserRole.USER);
        given(repo.findUserById(xxx.getUid())).willReturn(Optional.of(xxx));
        ;
        assertThat(test.loadUserByUsername("hoavq")).isNotNull();
    }
    @Test
    void checkLoadUserByUsernameNonExisted() {
        User hoavq = new User("hoavq", "1234",5, UserRole.USER);
        assertThatThrownBy(()->test.loadUserByUsername(hoavq.getUid())).isInstanceOf(UsernameNotFoundException.class);
    }
}