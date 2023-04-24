package com.manabie.todo.service.impl;

import com.manabie.todo.entity.UserEntity;
import com.manabie.todo.model.CreateUserRequest;
import com.manabie.todo.model.UserInfo;
import com.manabie.todo.repository.UserRepository;
import com.manabie.todo.service.UserService;
import lombok.RequiredArgsConstructor;
import org.modelmapper.ModelMapper;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements UserService {
    final UserRepository userRepository;
    final ModelMapper modelMapper;

    @Override
    public UserInfo create(CreateUserRequest request) {
        UserEntity userEntity = modelMapper.map(request, UserEntity.class);
        userEntity.setCreatedAt(LocalDateTime.now());
        return modelMapper.map(userRepository.save(userEntity), UserInfo.class);
    }

    @Override
    public UserInfo getById(Long userId) {
        return modelMapper.map(userRepository.getReferenceById(userId), UserInfo.class);
    }

    @Override
    public List<UserInfo> findAll() {
        return userRepository.findAll().stream().map(u -> modelMapper.map(u, UserInfo.class)).toList();
    }
}
