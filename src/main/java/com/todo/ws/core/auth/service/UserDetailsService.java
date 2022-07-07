package com.todo.ws.core.auth.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

import com.todo.ws.core.auth.dto.UserDetailsDto;
import com.todo.ws.core.auth.exception.UserDoesNotExistException;
import com.todo.ws.core.auth.model.User;
import com.todo.ws.core.auth.repository.UserRepository;

@Service
public class UserDetailsService  implements org.springframework.security.core.userdetails.UserDetailsService {

	@Autowired
    private  UserRepository userRepository;



    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        final User user = this.userRepository
            .findByUsername(username)
            .orElseThrow(() -> new UserDoesNotExistException("User not found with username: " + username));

        return new UserDetailsDto(user);
    }
}
