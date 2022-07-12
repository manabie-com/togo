package com.todo.ws.user.service;

import com.todo.ws.user.application.dto.CustomUserDetails;
import com.todo.ws.user.exception.UserDoesNotExistException;
import com.todo.ws.user.model.TodoUser;
import com.todo.ws.user.repository.TodoUserRepository;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

@Service
public class CustomUserDetailsService implements UserDetailsService {

    private final TodoUserRepository todoUserRepository;

    public CustomUserDetailsService(TodoUserRepository todoUserRepository) {
        this.todoUserRepository = todoUserRepository;
    }

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        final TodoUser todoUser = this.todoUserRepository
            .findByUsername(username)
            .orElseThrow(() -> new UserDoesNotExistException("User not found with username: " + username));

        return new CustomUserDetails(todoUser);
    }
}
