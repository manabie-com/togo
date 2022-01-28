package com.thang.todo.controller;

import javax.validation.Valid;

import com.thang.todo.config.WebSecurityConfig;
import com.thang.todo.entities.User;
import com.thang.todo.payload.LoginRequest;
import com.thang.todo.payload.LoginResponse;
import com.thang.todo.payload.RegisterRequest;
import com.thang.todo.repositories.UserRepository;
import com.thang.todo.services.UserDetailsImpl;
import com.thang.todo.utils.jwt.JwtTokenProvider;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@CrossOrigin
@RestController
@RequestMapping("/api")
public class UserController {

    private static Long DEFAULT_MAXIMUM_TASKS = 5L;

    @Autowired
    AuthenticationManager authenticationManager;

    @Autowired
    WebSecurityConfig webSecurityConfig;

    @Autowired
    private JwtTokenProvider tokenProvider;

    @Autowired
    UserRepository userRepository;

    @PostMapping("/login")
    public LoginResponse authenticateUser(@Valid @RequestBody LoginRequest loginRequest) {
        Authentication authentication = authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(
                        loginRequest.getUsername(),
                        loginRequest.getPassword()
                )
        );

        SecurityContextHolder.getContext().setAuthentication(authentication);

        String jwt = tokenProvider.generateToken((UserDetailsImpl) authentication.getPrincipal());
        return new LoginResponse(jwt);
    }

    @PostMapping("/register")
    public void registerUser(@Valid @RequestBody RegisterRequest registerRequest) {
        PasswordEncoder passwordEncoder = webSecurityConfig.passwordEncoder();
        String encodedPassword = passwordEncoder.encode(registerRequest.getPassword());
        User newUser = new User(registerRequest.getUsername(), encodedPassword, DEFAULT_MAXIMUM_TASKS);
        userRepository.save(newUser);
    }
}