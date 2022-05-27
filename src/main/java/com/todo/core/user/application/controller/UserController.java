package com.todo.core.user.application.controller;

import com.todo.core.commons.model.GenericResponse;
import com.todo.core.user.application.dto.UserLoginDTO;
import com.todo.core.user.application.dto.UserRegistrationDTO;
import com.todo.core.user.service.TodoUserService;
import com.todo.core.user.service.TokenProvider;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.config.annotation.authentication.configuration.AuthenticationConfiguration;
import org.springframework.security.core.Authentication;
import org.springframework.web.bind.annotation.*;

@RequestMapping("/api/v1/users")
@RestController
public class UserController {

    private final TodoUserService todoUserService;
    private final AuthenticationConfiguration authenticationConfiguration;
    private final TokenProvider tokenProvider;

    public UserController(TodoUserService todoUserService, AuthenticationConfiguration authenticationConfiguration,
                          TokenProvider tokenProvider) {
        this.todoUserService = todoUserService;
        this.authenticationConfiguration = authenticationConfiguration;
        this.tokenProvider = tokenProvider;
    }

    @PostMapping(value = "/create", consumes = "application/json")
    public GenericResponse<Boolean> createUser(@RequestBody UserRegistrationDTO dto) {
        return this.todoUserService.createUser(dto);
    }

    @PostMapping(value = "/login", consumes="application/json", produces = "application/json")
    public GenericResponse<String> authenticateUser(@RequestBody UserLoginDTO loginRequest) throws Exception {
        Authentication authentication = authenticationConfiguration.getAuthenticationManager().authenticate(
            new UsernamePasswordAuthenticationToken(
                loginRequest.getUsername(),
                loginRequest.getPassword()
            )
        );
//        SecurityContextHolder.getContext().setAuthentication(authentication);
        String jwt = tokenProvider.generateToken(authentication);

        return new GenericResponse<>(
            "SUCCESS", jwt);
    }
}
