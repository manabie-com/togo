package com.example.demo.model;

import com.example.demo.model.RegisterRequest;
import com.example.demo.model.AuthRequest;
import com.example.demo.model.AuthResponse;
import com.example.demo.model.MyAuthorities;
import com.example.demo.model.Task;
import com.example.demo.model.UpdateLimitRequest;
import com.example.demo.model.User;
import com.example.demo.model.UserSettings;
import com.example.demo.repository.AuthorityRepository;
import com.example.demo.repository.UserRepository;
import com.example.demo.repository.UserSettingsRepository;
import com.example.demo.service.UserService;
import com.example.demo.util.JwtTokenUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.util.Optional;

@RestController
@RequestMapping
public class AuthController {

    @Autowired
    @Qualifier("authenticationManager")
    private AuthenticationManager authenticationManager;

    @Autowired
    private JwtTokenUtil jwtTokenUtil;

    @Autowired
    private UserService userService;

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private UserSettingsRepository userSettingsRepository;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @Autowired
    private AuthorityRepository authorityRepository;

    @Autowired
    private MyAuthorities userAuth;

    @PostMapping("auth")
    public ResponseEntity auth(@RequestBody AuthRequest authRequest) throws Exception {
        authenticate(authRequest.getUsername(), authRequest.getPassword());
        User user = userService.loadUserByUsername(authRequest.getUsername());
        String token = jwtTokenUtil.generateToken(user);
        return ResponseEntity.ok(new AuthResponse(token));
    }

    @PutMapping("update-limit")
    public ResponseEntity updateLimit(HttpServletRequest request, @RequestBody UpdateLimitRequest updateLimitRequest) {
        final String requestTokenHeader = request.getHeader("Authorization");
        String jwtToken = requestTokenHeader.substring(7);
        User user = userService.loadUserByUsername(jwtTokenUtil.getUsernameFromToken(jwtToken));
        UserSettings userSettings = userSettingsRepository.findByUser(user).get();
        userSettings.setDailyLimit(updateLimitRequest.getLimit());
        userSettingsRepository.save(userSettings);
        return ResponseEntity.ok(userSettings);
    }
    @PostMapping("register")
    public ResponseEntity authRegister(@RequestBody RegisterRequest registerRequest) throws Exception {
        Optional<User> opUser = userRepository.findByUsername(registerRequest.getUsername());
        if (opUser.isPresent()) {
            throw new Exception("User already exists");
        } else {
            User user = new User(registerRequest.getUsername(), passwordEncoder.encode(registerRequest.getPassword()));
            UserSettings userSettings = new UserSettings(user, registerRequest.getLimit());
            user.addAuthority(userAuth);
            userRepository.save(user);
            userSettingsRepository.save(userSettings);
            return ResponseEntity.ok(user);
        }

    }

    private void authenticate(String username, String password) throws Exception {
        try {
            authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(username, password));
        } catch (Exception e) {
            e.printStackTrace();
            throw new Exception("Auth Failed");
        }
    }
}
