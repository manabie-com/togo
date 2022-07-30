package com.uuhnaut69.app.user.api;

import com.uuhnaut69.app.user.model.User;
import com.uuhnaut69.app.user.service.UserService;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author uuhnaut
 */
@RestController
@RequiredArgsConstructor
@RequestMapping("/users")
public class UserController {

  private final UserService userService;

  /**
   * Find all user
   *
   * @return Return {@link List} contain all {@link User}
   */
  @GetMapping
  public List<User> findAllUser() {
    return userService.findAllUser();
  }

  /**
   * Find user by user id
   *
   * @param userId User id must not be {@literal null}
   * @return Return {@link User}
   */
  @GetMapping("/{userId}")
  public User findUserDetail(@PathVariable Long userId) {
    return userService.findUserById(userId);
  }
}
