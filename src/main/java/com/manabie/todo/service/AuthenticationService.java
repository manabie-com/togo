package com.manabie.todo.service;

import com.manabie.todo.domain.Credential;
import com.manabie.todo.domain.User;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Slf4j
@Service
@AllArgsConstructor
public class AuthenticationService {

  private final UserService userService;
  private final PasswordEncoder passwordEncoder;

  public Optional<User> authentication(Credential credential) {
    var user = userService.getByUsername(credential.getUsername());

    if (user.isEmpty()) {
      return Optional.empty();
    }

    if (passwordEncoder.matches(credential.getPassword(), user.get().getPassword())) {
      return user;
    }

    return Optional.empty();
  }
}
