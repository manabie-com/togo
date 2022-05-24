package com.todo.core.user.service;

import com.todo.core.user.application.dto.CustomUserDetails;
import com.todo.core.user.exception.UserDoesNotExistException;
import com.todo.core.user.model.TodoUser;
import com.todo.core.user.repository.TodoUserRepository;
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
