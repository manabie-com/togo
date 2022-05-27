package com.todo.core.user.service;

import com.todo.core.commons.Messages;
import com.todo.core.commons.model.GenericResponse;
import com.todo.core.user.application.dto.UserRegistrationDTO;
import com.todo.core.user.exception.UserAlreadyExistsException;
import com.todo.core.user.model.TodoUser;
import com.todo.core.user.repository.TodoUserRepository;
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
    public GenericResponse<Boolean> createUser(UserRegistrationDTO userDto) {
        final String desiredUsername = userDto.getUsername();

        todoUserRepository.findByUsername(desiredUsername)
            .ifPresent((user) -> {throw new UserAlreadyExistsException("Username is already taken!");});

        userDto.doEncodePassword(this.passwordEncoder);

        final TodoUser userForSave = new TodoUser(userDto);
        todoUserRepository.save(userForSave);

        return new GenericResponse<>(true, Messages.USER_CREATE_SUCCESSFUL.getContent());
    }




}
