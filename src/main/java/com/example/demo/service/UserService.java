package com.example.demo.service;

import com.example.demo.model.User;
import com.example.demo.model.UserSettings;
import com.example.demo.repository.UserRepository;
import com.example.demo.repository.UserSettingsRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Service
public class UserService implements UserDetailsService {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private UserSettingsRepository userSettingsRepository;

    @Override
    public User loadUserByUsername(String username) throws UsernameNotFoundException {
        return userRepository.findById(username).get();
    }

    public UserSettings getUserSettings(User user) throws Exception {
        Optional<UserSettings> userSettingsOptional = userSettingsRepository.findByUser(user);
        if (userSettingsOptional.isPresent()) {
            return userSettingsOptional.get();
        }
        throw new Exception("User settings does not exist");
    }
}
