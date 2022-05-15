package com.manabie.todo.service;

import com.manabie.todo.domain.User;
import com.manabie.todo.entity.UserEntity;
import com.manabie.todo.exception.UserDuplicatedException;
import com.manabie.todo.mapper.UserMapper;
import com.manabie.todo.repository.UserRepository;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.context.annotation.Profile;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.test.context.ActiveProfiles;

import static org.junit.jupiter.api.Assertions.*;

@ActiveProfiles(profiles = "test")
@SpringBootTest
@EnableAutoConfiguration
class UserServiceTest {

  @Autowired
  UserService userService;

  @Autowired
  UserRepository userRepository;

  @Autowired
  PasswordEncoder passwordEncoder;

  @BeforeEach
  void setUp() {
  }

  @AfterEach
  void tearDown() {
    userRepository.deleteAll();
  }

  @Test
  void giveValidUser_whenAdd_thenSucceed() {
    var user = User.builder()
        .username("test")
        .password("test")
        .taskQuote(1L)
        .build();

    var result = userService.register(user);

    assertNotNull(result.getId());
  }

  @Test
  void giveExistUser_whenAdd_thenFailed() {
    var user = User.builder()
        .username("test")
        .password("test")
        .taskQuote(1L)
        .build();

    userService.register(user);
    assertThrows(UserDuplicatedException.class,()->{
      userService.register(user);
    });
  }

  @Test
  void giveExistUsername_whenFind_thenSucceed() {
    var user = UserEntity.builder()
        .username("test")
        .password("test")
        .taskQuote(1L)
        .build();

    userRepository.save(user);

    assertTrue(userService.getByUsername(user.getUsername()).isPresent());
  }

  @Test
  void giveNonExistUsername_whenFind_thenFailed() {
    var user = UserEntity.builder()
        .username("test")
        .password("test")
        .taskQuote(1L)
        .build();

    userRepository.save(user);

    assertFalse(userService.getByUsername(user.getUsername()+1).isPresent());
  }
  @Test
  void giveUser_whenUpdate_thenSucceed() {
    var userEntity = userRepository.save(UserEntity.builder()
        .username("test")
        .password(passwordEncoder.encode("test"))
        .taskQuote(1L)
        .build());

    var user = UserMapper.INSTANCE.toDto(userEntity);
    user.setTaskQuote(2L);

    assertEquals(userService.update(user).getTaskQuote(),user.getTaskQuote());
  }
}