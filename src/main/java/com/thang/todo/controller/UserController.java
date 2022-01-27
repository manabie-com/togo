package com.thang.todo.controller;

import javax.validation.Valid;

import com.thang.todo.config.WebSecurityConfig;
import com.thang.todo.entities.User;
import com.thang.todo.payload.LoginRequest;
import com.thang.todo.payload.LoginResponse;
import com.thang.todo.payload.RandomStuff;
import com.thang.todo.repositories.UserRepository;
import com.thang.todo.services.UserDetailsImpl;
import com.thang.todo.utils.jwt.JwtTokenProvider;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api")
public class UserController {

    @Autowired
    AuthenticationManager authenticationManager;

    @Autowired
    UserRepository userRepository;

    @Autowired
    WebSecurityConfig webSecurityConfig;

    @Autowired
    private JwtTokenProvider tokenProvider;

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
    public void registerUser(@Valid @RequestBody LoginRequest loginRequest) {
        PasswordEncoder passwordEncoder = webSecurityConfig.passwordEncoder();
        String encodedPassword = passwordEncoder.encode(loginRequest.getPassword());
        User newUser = new User(loginRequest.getUsername(), encodedPassword, 5L);
        userRepository.save(newUser);
    }

    @GetMapping("/random")
    public RandomStuff randomStuff(){
        return new RandomStuff("JWT Hợp lệ mới có thể thấy được message này");
    }

}