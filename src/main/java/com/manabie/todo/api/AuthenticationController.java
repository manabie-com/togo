package com.manabie.todo.api;

import com.manabie.todo.api.request.LoginRequest;
import com.manabie.todo.api.response.LoginResponse;
import com.manabie.todo.domain.Credential;
import com.manabie.todo.exception.UserUnauthorizedException;
import com.manabie.todo.service.AuthenticationService;
import com.manabie.todo.service.TokenService;
import com.manabie.todo.utils.AppResponse;
import lombok.AllArgsConstructor;
import org.springframework.http.MediaType;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@AllArgsConstructor
@RequestMapping(value = "/auth", consumes = MediaType.APPLICATION_JSON_VALUE)
public class AuthenticationController {

  private final AuthenticationService authenticationService;
  private final TokenService tokenService;


  @PostMapping("/login")
  public AppResponse<LoginResponse> login(@RequestBody @Validated LoginRequest request) {

    var user = authenticationService.authentication(Credential.builder()
        .username(request.getUsername())
        .password(request.getPassword())
        .build());

    if (user.isEmpty()) {
      throw new UserUnauthorizedException();
    }

    return AppResponse.ok(new LoginResponse(tokenService.generateToken(user.get())));
  }
}
