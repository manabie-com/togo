package com.manabie.todo.service;

import com.manabie.todo.model.CreateUserRequest;
import com.manabie.todo.model.UserInfo;

import java.util.List;

public interface UserService {
    UserInfo create(CreateUserRequest request);
    UserInfo getById(Long userId);
    List<UserInfo> findAll();
}
