package com.manabie.todo.service;

import com.manabie.todo.api.request.UpdateUserRequest;
import com.manabie.todo.configuration.AppProperties;
import com.manabie.todo.domain.User;
import com.manabie.todo.entity.UserEntity;
import com.manabie.todo.exception.UserDuplicatedException;
import com.manabie.todo.mapper.UserMapper;
import com.manabie.todo.repository.UserRepository;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Slf4j
@Service
@AllArgsConstructor
public class UserService {
  private final UserRepository userRepository;
  private final AppProperties appProperties;

  private final PasswordEncoder passwordEncoder;

  public User register(User user) {
    if (userRepository.existsByUsername(user.getUsername())) {
      throw new UserDuplicatedException();
    }
    log.debug("Add user: ",user);

    var userEntity = UserEntity.builder()
        .username(user.getUsername())
        .password(passwordEncoder.encode(user.getPassword()))
        .taskQuote(appProperties.getDefaultQuoteTask())
        .build();

    var saved = userRepository.saveAndFlush(userEntity);

    return UserMapper.INSTANCE.toDto(saved);
  }

  public Optional<User> getByUsername(String username) {
    var user = userRepository.findByUsername(username);
    if (user.isEmpty()) {
      return Optional.empty();
    }

    return Optional.of(UserMapper.INSTANCE.toDto(user.get()));
  }

  public User update(User updateUser) {
    log.debug("User: {}, update: {}",updateUser.getId(), updateUser);
    var userEntity = UserEntity.builder()
        .id(updateUser.getId())
        .username(updateUser.getUsername())
        .password(passwordEncoder.encode(updateUser.getPassword()))
        .taskQuote(updateUser.getTaskQuote())
        .build();

    var saved = userRepository.saveAndFlush(userEntity);

    return UserMapper.INSTANCE.toDto(saved);
  }
}
