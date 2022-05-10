package com.todo.controller;

import com.todo.config.JwtTokenUtil;
import com.todo.model.*;
import com.todo.repository.AppAccountRepository;
import com.todo.service.account.JwtUserDetailsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.DisabledException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;

/*
Expose a POST API /authenticate using the JwtAuthenticationController. The POST API gets username and password in the
body- Using Spring Authentication Manager we authenticate the username and password.If the credentials are valid,
a JWT token is created using the JWTTokenUtil and provided to the client.
 */
@RestController
@CrossOrigin
public class JwtAuthenticationController {

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private JwtTokenUtil jwtTokenUtil;

    @Autowired
    AppAccountRepository appAccountRepository;

    @Autowired
    private JwtUserDetailsService userDetailsService;

    // API for testing
    @GetMapping("/hello")
    public ResponseEntity<?> hello() {
        return ResponseEntity.ok(new MessageDTO("Hello"));
    }

    // register new user
    @RequestMapping(value = "/register", method = RequestMethod.POST)
    public ResponseEntity<?> saveUser(@RequestBody AccountDTO account, HttpServletRequest request) throws Exception {
        // prevent same email register again
        if (appAccountRepository.findByUsername(account.getUsername()) != null) {
            return ResponseEntity.ok(new MessageDTO("Email đã được đăng kí"));
        }

        String siteURL = request.getRequestURL().toString();
        return ResponseEntity.ok(userDetailsService.save(account));
    }

    // authenticate current user
    @RequestMapping(value = "/authenticate", method = RequestMethod.POST)
    public ResponseEntity<?> createAuthenticationToken(@RequestBody JwtRequest authenticationRequest) throws Exception {
        String username = authenticationRequest.getUsername();

        // check username exist
        if (appAccountRepository.findByUsername(username) == null) {
            return ResponseEntity.ok(new MessageDTO("Email chưa được đăng kí"));
        }

        authenticate(username, authenticationRequest.getPassword());

        final UserDetails userDetails = userDetailsService.loadUserByUsername(username);

        final String token = jwtTokenUtil.generateToken(userDetails);

        Long id = appAccountRepository.findByUsername(username).getId();
        // find information to return based on account role

        return ResponseEntity.ok(new JwtResponse(id, username, token));
    }

    // method that do the authentication process
    private void authenticate(String username, String password) throws Exception {
        try {
            authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(username, password));
        } catch (DisabledException e) {
            throw new Exception("USER_DISABLED", e);
        } catch (BadCredentialsException e) {
            throw new Exception("INVALID_CREDENTIALS", e);
        }
    }
}
