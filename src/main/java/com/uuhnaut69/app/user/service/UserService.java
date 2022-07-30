package com.uuhnaut69.app.user.service;

import com.uuhnaut69.app.common.exception.NotFoundException;
import com.uuhnaut69.app.user.model.User;
import com.uuhnaut69.app.user.repository.UserRepository;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.lang.NonNull;
import org.springframework.stereotype.Service;

/**
 * @author uuhnaut
 */
@Service
@RequiredArgsConstructor
public class UserService {

  private final UserRepository userRepository;

  /**
   * Get all user
   *
   * @return Return a {@link List} contain {@link User}
   */
  public List<User> findAllUser() {
    return userRepository.findAll();
  }

  /**
   * Find user by user id
   *
   * @param userId User id must not be {@literal null}
   * @return Return {@link User}
   */
  public User findUserById(@NonNull Long userId) {
    return userRepository.findById(userId)
        .orElseThrow(() -> new NotFoundException(String.format("User %s not found", userId)));
  }
}
