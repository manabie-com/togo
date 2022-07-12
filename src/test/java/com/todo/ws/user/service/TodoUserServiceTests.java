package com.todo.ws.user.service;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import com.todo.ws.user.application.dto.UserRegistrationDTO;
import com.todo.ws.user.exception.UserAlreadyExistsException;
import com.todo.ws.user.model.TodoUser;
import com.todo.ws.user.repository.TodoUserRepository;

import java.util.Optional;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
public class TodoUserServiceTests {

    @InjectMocks
    TodoUserService todoUserService;

    @Mock
    TodoUserRepository todoUserRepository;

    @Test
    public void testCreateUserSuccess() {
        final UserRegistrationDTO userRegistrationDTO = mock(UserRegistrationDTO.class);
        when(userRegistrationDTO.getUsername()).thenReturn("testUser");
        when(this.todoUserRepository.findByUsername("testUser")).thenReturn(Optional.empty());

        assertThat(todoUserService.createUser(userRegistrationDTO).getData()).isTrue();
    }

    @Test
    public void testCreateUserExistsException() {
        final UserRegistrationDTO userRegistrationDTO = mock(UserRegistrationDTO.class);
        when(userRegistrationDTO.getUsername()).thenReturn("testUser");
        when(this.todoUserRepository.findByUsername("testUser")).thenReturn(Optional.of(mock(TodoUser.class)));

        Assertions.assertThrows(UserAlreadyExistsException.class,() -> todoUserService.createUser(userRegistrationDTO));
    }

}
