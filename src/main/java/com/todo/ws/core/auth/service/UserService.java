package com.todo.ws.core.auth.service;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.todo.ws.common.enums.AuthMessageEnum;
import com.todo.ws.common.model.ResponseEntity;
import com.todo.ws.core.auth.dto.RegRequestDto;
import com.todo.ws.core.auth.exception.UserAlreadyExistsException;
import com.todo.ws.core.auth.model.User;
import com.todo.ws.core.auth.repository.UserRepository;

@Service
@Transactional(readOnly = false)
public class UserService {
	
	@Autowired
    private UserRepository todoUserRepository;
	@Autowired
    private PasswordEncoder passwordEncoder;


    @Transactional(readOnly = false)
    public ResponseEntity<Boolean> createUser(RegRequestDto userDto) {
        final String desiredUsername = userDto.getUsername();

        todoUserRepository.findByUsername(desiredUsername)
            .ifPresent((user) -> {throw new UserAlreadyExistsException("Username is already taken!");});

        userDto.doEncodePassword(this.passwordEncoder);

        final User userForSave = new User(userDto);
        todoUserRepository.save(userForSave);

        return new ResponseEntity<>(true, AuthMessageEnum.USER_CREATE_SUCCESSFUL.getContent());
    }
}
