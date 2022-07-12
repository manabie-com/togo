package com.todo.ws.user.service;

import com.todo.ws.commons.enums.AuthenticationEnum;
import com.todo.ws.commons.model.ResponseEntity;
import com.todo.ws.user.application.dto.UserRegistrationDTO;
import com.todo.ws.user.exception.UserAlreadyExistsException;
import com.todo.ws.user.model.TodoUser;
import com.todo.ws.user.repository.TodoUserRepository;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@Transactional(readOnly = false)
public class TodoUserService implements UserService {

    private final TodoUserRepository todoUserRepository;
    private final PasswordEncoder passwordEncoder;

    public TodoUserService(TodoUserRepository todoUserRepository,
                           PasswordEncoder passwordEncoder) {
        this.todoUserRepository = todoUserRepository;
        this.passwordEncoder = passwordEncoder;
    }

    @Transactional(readOnly = false)
    public ResponseEntity<Boolean> createUser(UserRegistrationDTO userDto) {
        final String desiredUsername = userDto.getUsername();

        todoUserRepository.findByUsername(desiredUsername)
            .ifPresent((user) -> {throw new UserAlreadyExistsException("Username is already taken!");});

        userDto.doEncodePassword(this.passwordEncoder);

        final TodoUser userForSave = new TodoUser(userDto);
        todoUserRepository.save(userForSave);

        return new ResponseEntity<>(true, AuthenticationEnum.USER_CREATE_SUCCESSFUL.getContent());
    }




}
