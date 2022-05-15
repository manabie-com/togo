package com.manabie.todo.api;

import com.manabie.todo.api.request.RegisterRequest;
import com.manabie.todo.api.request.UpdateUserRequest;
import com.manabie.todo.api.response.UserResponse;
import com.manabie.todo.domain.User;
import com.manabie.todo.service.UserService;
import com.manabie.todo.utils.AppResponse;
import lombok.AllArgsConstructor;
import org.springframework.http.MediaType;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

@RestController
@AllArgsConstructor
@RequestMapping(value = "/user", consumes = MediaType.APPLICATION_JSON_VALUE)
public class UserController {

  private final UserService userService;

  @PostMapping("/register")
  public AppResponse<UserResponse> register(@RequestBody @Validated RegisterRequest request) {

    var user = userService.register(
        User.builder()
            .username(request.getUsername())
            .password(request.getPassword())
            .build());

    return AppResponse.ok(
        UserResponse.builder()
            .id(user.getId())
            .username(user.getUsername())
            .taskQuote(user.getTaskQuote())
            .build()
    );
  }

  @PutMapping()
  public AppResponse<UserResponse> update(@RequestBody UpdateUserRequest request){
    var user = (User)SecurityContextHolder.getContext().getAuthentication().getPrincipal();

    var updateUser = userService.update(User.builder()
        .username(request.getUsername() != null ? request.getUsername() : user.getUsername())
        .password(request.getPassword() != null ? request.getPassword() : user.getPassword())
        .taskQuote(request.getTaskQuote() != null ? request.getTaskQuote() : user.getTaskQuote())
        .build());

    return AppResponse.ok(
        UserResponse.builder()
            .id(updateUser.getId())
            .username(updateUser.getUsername())
            .taskQuote(updateUser.getTaskQuote())
            .build()
    );
  }
}
