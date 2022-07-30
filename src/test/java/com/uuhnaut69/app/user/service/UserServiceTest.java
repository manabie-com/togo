package com.uuhnaut69.app.user.service;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import com.uuhnaut69.app.common.exception.NotFoundException;
import com.uuhnaut69.app.user.model.User;
import com.uuhnaut69.app.user.repository.UserRepository;
import java.time.Instant;
import java.util.List;
import java.util.Optional;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.junit.jupiter.MockitoExtension;

/**
 * @author uuhnaut
 */
@ExtendWith(MockitoExtension.class)
class UserServiceTest {

  @Mock
  private UserRepository userRepository;

  @InjectMocks
  private UserService userService;

  @Test
  void findAllUser() {
    var users = List.of(new User(1L, "uuhnaut69", 10, Instant.now()),
        new User(2L, "uuhnaut96", 69, Instant.now()));

    Mockito.when(userRepository.findAll()).thenReturn(users);

    var result = userService.findAllUser();

    assertEquals(2, result.size());
  }

  @Test
  void findUserByIdSuccessful() {
    var user = new User(2L, "uuhnaut96", 69, Instant.now());

    Mockito.when(userRepository.findById(2L)).thenReturn(Optional.of(user));

    var result = userService.findUserById(user.getId());

    assertEquals(2L, result.getId());
  }

  @Test
  void findUserByIdNotFound() {
    Mockito.when(userRepository.findById(1L)).thenReturn(Optional.empty());

    try {
      userService.findUserById(1L);
    } catch (Exception e) {
      assertTrue(e instanceof NotFoundException);
    }
  }
}